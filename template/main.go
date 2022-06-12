package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/caarlos0/env"
	"github.com/razielsd/rzgrpcmock/template/internal/app"
	"github.com/razielsd/rzgrpcmock/template/internal/config"
	"go.uber.org/zap"
)

func main() {
	cfg := &config.Config{}

	if err := env.Parse(cfg); err != nil {
		zap.S().Fatalf("failed to retrieve env variables, %v", err)
	}
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	if err := app.Run(ctx, cfg); err != nil {
		zap.S().Fatal("error running grpc server ", err)
	}
	defer cancel()
}
