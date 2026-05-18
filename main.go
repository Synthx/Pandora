package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"pandora/login"
	"syscall"

	"go.uber.org/zap"
)

var (
	port = flag.Int("port", 443, "port to listen on")
)

func initLogger() *zap.Logger {
	if os.Getenv("APP_ENV") == "development" {
		return zap.Must(zap.NewDevelopment())
	}

	return zap.Must(zap.NewProduction())
}

func main() {
	flag.Parse()

	logger := initLogger()
	defer logger.Sync()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	loginServer, err := login.NewServer(*port, logger)
	if err != nil {
		logger.Error("Failed to create server", zap.Error(err))
		os.Exit(1)
	}

	if err = loginServer.Serve(ctx); err != nil {
		logger.Error("Failed to start server", zap.Error(err))
		os.Exit(1)
	}
}
