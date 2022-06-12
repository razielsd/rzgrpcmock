package interceptor

import (
	"context"
	"github.com/razielsd/rzgrpcmock/template/internal/reqmatcher"
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
	meta := reqmatcher.RequestMeta{
		Method: method,
	}
	ctx = context.WithValue(ctx, reqmatcher.MetaKey, meta)
	reply, err := handler(ctx, req)
	return reply, err
}
