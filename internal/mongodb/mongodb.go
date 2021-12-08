package mongodb

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	Client      *mongo.Client
	Ctx         context.Context
	Collections map[string]*mongo.Collection
}

func NewMongoClient() (*MongoClient, error) {
	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, errors.New("error creating mongodb client")
	}

	ep := client.Ping(ctx, nil)
	if ep != nil {
		return nil, errors.New("error connecting to mongodb")
	}

	collections := map[string]*mongo.Collection{
		"uplay": client.Database("gor6").Collection("uplay"),
		"psn":   client.Database("gor6").Collection("psn"),
		"xbl":   client.Database("gor6").Collection("xbl"),
	}

	return &MongoClient{
		Client:      client,
		Ctx:         ctx,
		Collections: collections,
	}, nil
}

func (c *MongoClient) Close() {
	c.Client.Disconnect(c.Ctx)
}
