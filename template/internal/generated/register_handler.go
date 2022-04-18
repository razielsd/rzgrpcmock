package generated

import (
	"sync"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	handlerRegistration []func(s *grpc.Server, logger *zap.Logger)
	mu                  = &sync.Mutex{} //nolint:deadcode,unused,varcheck
)

func RegisterHandler(s *grpc.Server, logger *zap.Logger) {
	for _, f := range handlerRegistration {
		f(s, logger)
	}
}
