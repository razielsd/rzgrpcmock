package generated

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	handlerRegistration []func(s *grpc.Server, logger *zap.Logger)
)

func RegisterHandler(s *grpc.Server, logger *zap.Logger) {
	for _, f := range handlerRegistration {
		f(s, logger)
	}
}
