package commons

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
)

func CreateGRPCServer(grpcAddr string) *grpc.Server {
	grpcPanicRecoveryHandler := func(p any) (err error) {
		zap.L().Error("recovered from panic", zap.Any("panic", p))
		return status.Errorf(codes.Internal, "internal server error")
	}

	// Create recovery options with custom handler
	recoveryOpts := []recovery.Option{
		recovery.WithRecoveryHandler(grpcPanicRecoveryHandler),
	}

	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			recovery.UnaryServerInterceptor(recoveryOpts...),
		),
		grpc.ChainStreamInterceptor(
			recovery.StreamServerInterceptor(recoveryOpts...),
		),
	)

	return s
}
