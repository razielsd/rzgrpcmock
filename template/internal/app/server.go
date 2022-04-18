package app

import (
	"context"
	"net"
	"time"

	"github.com/razielsd/rzgrpcmock/server/internal/mock"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/razielsd/rzgrpcmock/server/internal/config"
	"github.com/razielsd/rzgrpcmock/server/internal/generated"
	"github.com/razielsd/rzgrpcmock/server/internal/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
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
	lg, err := logger.GetLogger(cfg)
	if err != nil {
		zap.S().Fatalf("Unable init logger: %s", err)
	}
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
		zap.S().Fatalf("failed to listen tcp %s, %v\n", cfg.GRPCAddr, err)
	}

	generated.RegisterHandler(s, lg)

	go func() {
		lg.Info("starting grpc server", zap.String("grpc host", cfg.GRPCAddr))
		if err := s.Serve(l); err != nil {
			zap.S().Fatalf("error service grpc server %v", err)
		}
	}()

	apiServer := mock.NewApiServer(cfg, lg)
	apiServer.Run(ctx)

	gracefulShutDown(ctx, s, cancel)

	return nil
}

// Will be generated same code
//  func registerHandler(s *grpc.Server, logger *log.Logger) {
//	// grpc handlers
//	myService := myservice.NewService(logger)
//
//	// register handlers
//	papi.RegisterProjectServer(s, myService)
//
// }

func gracefulShutDown(ctx context.Context, s *grpc.Server, cancel context.CancelFunc) {
	<-ctx.Done()
	zap.S().Info("Received shutdown signal, graceful shutdown done")
	s.GracefulStop()
	cancel()
}
