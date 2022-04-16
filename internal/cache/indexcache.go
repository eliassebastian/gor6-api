package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type IndexCache struct {
	DB *redis.Client
}

func InitIndexCache(ctx context.Context) (*IndexCache, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // use default Addr
		Password: "",               // no password set
		DB:       0,                // use default DB
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return &IndexCache{
		DB: client,
	}, err
}

func (c *IndexCache) Close() {
	c.DB.Close()
}
