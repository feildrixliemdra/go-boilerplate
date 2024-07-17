package bootstrap

import (
	"context"
	"go-boilerplate/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitiateMongoDB(cfg *config.Config) (*mongo.Client, error) {
	client, err := mongo.Connect(context.Background(), options.Client().
		ApplyURI(cfg.MongoDB.URL))
	if err != nil {
		return nil, err
	}

	return client, nil
}
