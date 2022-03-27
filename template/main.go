package main

import (
	"github.com/razielsd/rzgrpcmock/server/internal/app"
	"github.com/razielsd/rzgrpcmock/server/internal/config"
	"os"
	"os/signal"

	"context"
	"github.com/caarlos0/env"
	log "github.com/sirupsen/logrus"
)

func main() {
	cfg := &config.Config{}

	if err := env.Parse(cfg); err != nil {
		log.Fatalf("failed to retrieve env variables, %v", err)
	}
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	if err := app.Run(ctx, cfg); err != nil {
		log.Fatal("error running grpc server ", err)
	}
	defer cancel()
}
