package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"helicopter-hr/config"
)

func InitMongoConnection(cfg *config.ConfigApp) (db *mongo.Client, err error) {
	connectionString := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",
		cfg.MongoDB.Username,
		cfg.MongoDB.Password,
		cfg.MongoDB.HostName,
		cfg.MongoDB.Port,
		cfg.MongoDB.DatabaseName)

	mongoClient, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(connectionString),
	)

	if err != nil {
		return
	}

	err = mongoClient.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		return
	}

	return mongoClient, nil

}
