package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Account struct {
	Id          bson.ObjectID `bson:"_id"`
	Username    string
	Token       string
	IsAdmin     bool
	IsBanned    bool
	BannedUntil time.Time
}

type AccountRepository struct {
	collection *mongo.Collection
}

func NewAccountRepository(database *mongo.Database) *AccountRepository {
	return &AccountRepository{collection: database.Collection("accounts")}
}

func (r *AccountRepository) GetByToken(ctx context.Context, id string) (*Account, error) {
	filter := bson.M{"token": id}

	var result Account
	err := r.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
