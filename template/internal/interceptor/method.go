package interceptor

import (
	"context"
	"strings"

	"google.golang.org/grpc"
)

func UnaryMethodServerInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	method := strings.ToLower(strings.TrimLeft(info.FullMethod, `/`))
	ctx = context.WithValue(ctx, "method", method) //nolint: revive, staticcheck
	reply, err := handler(ctx, req)
	return reply, err
}
