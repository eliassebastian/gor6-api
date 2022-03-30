package mongodb

import (
	"context"
	"errors"
	model "github.com/eliassebastian/gor6-api/cmd/api/graph/models"
	"go.mongodb.org/mongo-driver/bson"
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
		return nil, errors.New(ep.Error())
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

func (c *MongoClient) SearchPlayers(ctx context.Context, p, n string) ([]*model.PlayerSearchResults, error) {
	//text search for name
	filter := bson.M{
		"$text": bson.M{
			"$search": n,
		},
	}
	//Limit returns to 25 results
	opts := options.Find().SetLimit(25)
	//projection - fields to return from search (_id included)
	opts.SetProjection(bson.D{
		{"item", 1},
		{"status", 1},
	})
	//sort
	opts.SetSort(
		bson.M{
			"score": bson.M{"$meta": "textScore"},
		},
	)
	//run query
	cursor, err := c.Collections[p].Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	//decode into slice of players
	var articles []*model.PlayerSearchResults
	de := cursor.All(ctx, articles)
	if de != nil {
		return nil, de
	}

	return articles, nil
}

func (c *MongoClient) NewPlayer(ctx context.Context, p string, document *model.Player) error {
	_, err := c.Collections[p].InsertOne(ctx, document)
	return err
}

func (c *MongoClient) Player(ctx context.Context, p, id string) (*model.Player, error) {
	var player *model.Player

	err := c.Collections[p].FindOne(ctx, bson.D{{"_id", id}}).Decode(&player)
	if err != nil {
		return nil, err
	}

	return player, nil
}
