package login

import (
	"context"
	"pandora/internal/login/repository"
	"pandora/internal/pkg"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

func NewContainer(ctx context.Context) (*dig.Container, error) {
	container := dig.New()

	err := container.Provide(func() context.Context {
		return ctx
	})
	if err != nil {
		return nil, err
	}

	err = container.Provide(NewConfig)
	if err != nil {
		return nil, err
	}

	err = container.Provide(func(c *Config) (*zap.Logger, error) {
		return pkg.NewLogger(c.Env, c.LogLevel)
	})
	if err != nil {
		return nil, err
	}

	err = container.Provide(func(ctx context.Context, c *Config, logger *zap.Logger) (*mongo.Database, error) {
		return pkg.NewDatabase(ctx, c.DatabaseUri, c.DatabaseName, logger)
	})
	if err != nil {
		return nil, err
	}

	err = container.Provide(repository.NewAccountRepository)
	if err != nil {
		return nil, err
	}

	err = container.Provide(NewSessionHandler)
	if err != nil {
		return nil, err
	}

	err = container.Provide(NewMessageDispatcher)
	if err != nil {
		return nil, err
	}

	err = container.Provide(func(c *Config, logger *zap.Logger, dispatcher *MessageDispatcher) (*Server, error) {
		return NewServer(c.ServerPort, logger, dispatcher)
	})
	if err != nil {
		return nil, err
	}

	return container, nil
}
