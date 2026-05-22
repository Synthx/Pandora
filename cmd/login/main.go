package main

import (
	"context"
	"os"
	"os/signal"
	"pandora/internal/login"
	"syscall"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	container, err := login.NewContainer(ctx)
	if err != nil {
		panic(err)
	}

	err = container.Invoke(func(server *login.Server, logger *zap.Logger, db *mongo.Database) error {
		defer func() {
			if err := logger.Sync(); err != nil {
				logger.Error("Failed to sync logger", zap.Error(err))
			}
		}()
		defer func() {
			if err := db.Client().Disconnect(ctx); err != nil {
				logger.Error("Failed to disconnect from database", zap.Error(err))
			}
		}()

		return server.Serve(ctx)
	})
	if err != nil {
		panic(err)
	}
}
