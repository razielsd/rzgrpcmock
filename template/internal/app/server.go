package app

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/razielsd/rzgrpcmock/template/internal/interceptor"

	"github.com/razielsd/rzgrpcmock/template/internal/mockserver"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/razielsd/rzgrpcmock/template/internal/config"
	"github.com/razielsd/rzgrpcmock/template/internal/generated"
	"github.com/razielsd/rzgrpcmock/template/internal/logger"
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
	lg, err := logger.CreateLogger(cfg)
	if err != nil {
		log.Fatalln("unable init logger")
	}
	grpc_prometheus.EnableHandlingTimeHistogram()
	// grpc server
	s := grpc.NewServer(
		grpc.KeepaliveParams(keepAliveParams),
		grpc.ChainUnaryInterceptor(
			grpc_prometheus.UnaryServerInterceptor,
			interceptor.UnaryMethodServerInterceptor,
		),
	)
	_, cancel := context.WithCancel(ctx)

	l, err := net.Listen("tcp", cfg.GRPCAddr)
	if err != nil {
		lg.Fatal("failed to listen tcp", zap.Error(err), zap.String("port", cfg.GRPCAddr))
	}

	generated.RegisterHandler(s, lg)

	go func() {
		lg.Info("starting grpc server", zap.String("grpc host", cfg.GRPCAddr))
		if err := s.Serve(l); err != nil {
			lg.Fatal("error service grpc server", zap.Error(err))
		}
	}()

	apiServer := mockserver.NewApiServer(cfg, lg)
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
