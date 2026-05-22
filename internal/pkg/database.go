package pkg

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.uber.org/zap"
)

func NewDatabase(ctx context.Context, uri string, name string, logger *zap.Logger) (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(uri)
	clientOptions.BSONOptions = &options.BSONOptions{OmitEmpty: true}

	client, err := mongo.Connect(clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	logger.Info("Connected to database")

	return client.Database(name), nil
}
