package commons

import (
	"context"
	"time"

	recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/timeout"
	"github.com/mediaprodcast/commons/discovery"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func interceptorLogger(l *zap.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		zapFields := make([]zapcore.Field, len(fields)/2)
		for i := 0; i < len(fields); i += 2 {
			key, ok := fields[i].(string)
			if !ok {
				continue
			}
			value := fields[i+1]
			zapFields[i/2] = zap.Any(key, value)
		}
		switch lvl {
		case logging.LevelDebug:
			l.Debug(msg, zapFields...)
		case logging.LevelInfo:
			l.Info(msg, zapFields...)
		case logging.LevelWarn:
			l.Warn(msg, zapFields...)
		case logging.LevelError:
			l.Error(msg, zapFields...)
		default:
			l.Info(msg, zapFields...)
		}
	})
}

func NewGRPCServer(grpcAddr string) *grpc.Server {
	rpcLogger := zap.L().With(zap.String("service", "gRPC/server"))
	logTraceID := func(ctx context.Context) logging.Fields {
		if span := trace.SpanContextFromContext(ctx); span.IsSampled() {
			return logging.Fields{"traceID", span.TraceID().String()}
		}
		return nil
	}

	grpcPanicRecoveryHandler := func(p any) (err error) {
		zap.L().Error("recovered from panic", zap.Any("panic", p))
		return status.Errorf(codes.Internal, "internal server error")
	}

	// Create recovery options with custom handler
	recoveryOpts := []recovery.Option{
		recovery.WithRecoveryHandler(grpcPanicRecoveryHandler),
	}

	s := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
		grpc.ChainUnaryInterceptor(
			recovery.UnaryServerInterceptor(recoveryOpts...),
			logging.UnaryServerInterceptor(interceptorLogger(rpcLogger), logging.WithFieldsFromContext(logTraceID)),
		),
		grpc.ChainStreamInterceptor(
			recovery.StreamServerInterceptor(recoveryOpts...),
			logging.StreamServerInterceptor(interceptorLogger(rpcLogger), logging.WithFieldsFromContext(logTraceID)),
		),
	)

	return s
}

func NewGRPCClient(svcAddr string) (*grpc.ClientConn, error) {
	rpcLogger := zap.L().With(zap.String("service", "gRPC/client"))
	logTraceID := func(ctx context.Context) logging.Fields {
		if span := trace.SpanContextFromContext(ctx); span.IsSampled() {
			return logging.Fields{"traceID", span.TraceID().String()}
		}
		return nil
	}

	return grpc.NewClient(
		svcAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
		grpc.WithChainUnaryInterceptor(
			timeout.UnaryClientInterceptor(30*time.Second),
			logging.UnaryClientInterceptor(interceptorLogger(rpcLogger), logging.WithFieldsFromContext(logTraceID)),
		),
		grpc.WithChainStreamInterceptor(
			logging.StreamClientInterceptor(interceptorLogger(rpcLogger), logging.WithFieldsFromContext(logTraceID)),
		),
	)
}

func DiscoverNewGRPCClient(ctx context.Context, serviceName string, registry discovery.Registry) (*grpc.ClientConn, error) {
	svcAddr, err := discovery.Discover(ctx, serviceName, registry)
	if err != nil {
		return nil, err
	}

	return NewGRPCClient(svcAddr)
}
