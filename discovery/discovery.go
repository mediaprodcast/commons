package discovery

import (
	"context"
	"fmt"
	"math/rand"
	"time"
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
