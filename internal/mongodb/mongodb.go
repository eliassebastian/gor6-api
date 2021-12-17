package mongodb

import (
	"context"
	"errors"
	"fmt"
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
		"uplay": client.Database("gor6").Collection("pc"),
		"psn":   client.Database("gor6").Collection("ps4"),
		"xbl":   client.Database("gor6").Collection("xbox"),
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

func (c *MongoClient) NewPlayer(ctx context.Context, p string, document interface{}) interface{} {
	insert, err := c.Collections[p].InsertOne(ctx, document)
	if err != nil {
		fmt.Println("error inserting document into mongo:", err)
		return nil
	}

	return insert.InsertedID
}
