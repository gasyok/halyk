package interceptor

import (
	"context"
	"strings"

	"github.com/gasyok/Assessment/internal/pkg/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func WithLogging() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		resp, err = handler(ctx, req)
		_, method := splitFullMethodName(info.FullMethod)
		args := append(make([]zapcore.Field, 0, 2), zap.String("method", method))
		if s, ok := status.FromError(err); !ok {
			args = append(args, zap.Error(err))
		} else {
			args = append(args, zap.Any("code", s.Code()))
		}
		logger.Info("Request served", args...)
		return
	}
}

func splitFullMethodName(fullMethod string) (string, string) {
	fullMethod = strings.TrimPrefix(fullMethod, "/")
	if i := strings.Index(fullMethod, "/"); i >= 0 {
		return fullMethod[:i], fullMethod[i+1:]
	}
	return "unknown", "unknown"
}
