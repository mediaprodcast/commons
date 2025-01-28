package discovery

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/mediaprodcast/commons/discovery/consul"
	"go.uber.org/zap"
)

const (
	// Service Names
	IngressSvsName       = "ingress"
	TranscoderSvsName    = "transcoder"
	PackagerSvsName      = "packager"
	StorageSvsName       = "storage"
	OrchestrationSvsName = "orchestration"
	MonitorSvsName       = "monitor"
	APISvsName           = "api"
	DashboardSvsName     = "dashboard"
	ProbeSvsName         = "probe"
	DRMSvsName           = "drm"
	AutoscalerSvsName    = "autoscaler"
)

type Registry interface {
	Register(ctx context.Context, instanceID, serverName, hostPort string) error
	Deregister(ctx context.Context, instanceID, serviceName string) error
	Discover(ctx context.Context, serviceName string) ([]string, error)
	HealthCheck(instanceID, serviceName string) error
}

func GenerateInstanceID(serviceName string) string {
	return fmt.Sprintf("%s-%d", serviceName, rand.New(rand.NewSource(time.Now().UnixNano())).Int())
}

// Register sets up the service registry, starts health checks, and returns the instance ID and registry for later use
func Register(ctx context.Context, serviceName, serviceAddr string) (string, Registry, error) {
	registry, err := consul.NewRegistry(serviceName)
	if err != nil {
		return "", nil, err
	}

	instanceID := GenerateInstanceID(serviceName)
	if err := registry.Register(ctx, instanceID, serviceName, serviceAddr); err != nil {
		return "", nil, err
	}

	// Start health check goroutine
	go func() {
		ticker := time.NewTicker(time.Second * 1)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				if err := registry.HealthCheck(instanceID, serviceName); err != nil {
					zap.L().Error("Service failed health check", zap.String("service", serviceName), zap.Error(err))
				}
			case <-ctx.Done():
				zap.L().Info("Stopping health checks", zap.String("service", serviceName))
				return
			}
		}
	}()

	return instanceID, registry, nil
}

// Discover service address from registry
func Discover(ctx context.Context, serviceName string, registry Registry) (string, error) {
	addrs, err := registry.Discover(ctx, serviceName)
	if err != nil {
		return "", err
	}

	if len(addrs) == 0 {
		return "", fmt.Errorf("failed to discovered instances of %s", serviceName)
	}

	// Randomly select an instance
	svcAddr := addrs[rand.Intn(len(addrs))]

	zap.L().Info("Discovered instances of service",
		zap.Int("count", len(addrs)),
		zap.String("service", serviceName),
		zap.String("address", svcAddr),
	)

	return svcAddr, nil
}
