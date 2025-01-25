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
func Register(ctx context.Context, consulAddr, serviceName, serviceAddr string, l *zap.Logger) (string, Registry, error) {
	registry, err := consul.NewRegistry(consulAddr, serviceName)
	if err != nil {
		return "", nil, err
	}

	instanceID := GenerateInstanceID(serviceName)
	if err := registry.Register(ctx, instanceID, serviceName, serviceAddr); err != nil {
		return "", nil, err
	}

	// Start health check goroutine
	go func() {
		for {
			if err := registry.HealthCheck(instanceID, serviceName); err != nil {
				l.Error("Service failed health check", zap.String("service", serviceName), zap.Error(err))
			}
			time.Sleep(time.Second * 1)
		}
	}()

	return instanceID, registry, nil
}
