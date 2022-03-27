package generated

import (
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"sync"
)

var handlerRegistration []func(s *grpc.Server, logger *log.Logger)
var mu = &sync.Mutex{}

func RegisterHandler(s *grpc.Server, logger *log.Logger) {
	for _, f := range handlerRegistration {
		f(s, logger)
	}
}
