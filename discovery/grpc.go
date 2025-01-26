package discovery

import (
	"context"
	"fmt"
	"math/rand"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ServiceConnection(ctx context.Context, serviceName string, registry Registry) (*grpc.ClientConn, error) {
	addrs, err := registry.Discover(ctx, serviceName)
	if err != nil {
		return nil, err
	}

	if len(addrs) == 0 {
		return nil, fmt.Errorf("failed to discovered instances of %s", serviceName)
	}

	// Randomly select an instance
	svcAddr := addrs[rand.Intn(len(addrs))]

	zap.L().Info("Discovered instances of service",
		zap.Int("count", len(addrs)),
		zap.String("service", serviceName),
		zap.String("address", svcAddr),
	)

	return grpc.NewClient(
		svcAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// Add OpenTelemetry interceptors
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)
}
