package app

import (
	"context"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/razielsd/rzgrpcmock/server/internal/config"
	"github.com/razielsd/rzgrpcmock/server/internal/generated"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"net"
	"time"
)

var keepAliveParams = keepalive.ServerParameters{
	MaxConnectionIdle:     10 * time.Hour,
	MaxConnectionAge:      24 * time.Hour,
	MaxConnectionAgeGrace: 5 * time.Minute,
	Time:                  60 * time.Second,
	Timeout:               1 * time.Second,
}

func Run(ctx context.Context, cfg *config.Config) error {
	// log
	logger := log.StandardLogger()

	grpc_prometheus.EnableHandlingTimeHistogram()
	// grpc server
	s := grpc.NewServer(
		grpc.KeepaliveParams(keepAliveParams),
		grpc.ChainUnaryInterceptor(
			grpc_prometheus.UnaryServerInterceptor,
		),
	)
	_, cancel := context.WithCancel(ctx)

	l, err := net.Listen("tcp", cfg.GRPCAddr)
	if err != nil {
		log.Fatalf("failed to listen tcp %s, %v", cfg.GRPCAddr, err)
	}

	generated.RegisterHandler(s, logger)

	go func() {
		log.Infof("starting listening grpc server at %s", cfg.GRPCAddr)
		if err := s.Serve(l); err != nil {
			log.Fatalf("error service grpc server %v", err)
		}
	}()

	gracefulShutDown(s, ctx, cancel)

	return nil
}

// Will be generated same code
//func registerHandler(s *grpc.Server, logger *log.Logger) {
//	// grpc handlers
//	myService := myservice.NewService(logger)
//
//	// register handlers
//	papi.RegisterProjectServer(s, myService)
//
//}

func gracefulShutDown(s *grpc.Server, ctx context.Context, cancel context.CancelFunc) {
	<-ctx.Done()
	errorMessage := "Received shutdown signal, graceful shutdown done"
	log.Info(errorMessage)
	s.GracefulStop()
	cancel()
}
