package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type ProfileCache struct {
	DB *redis.Client
}

func InitProfileCache(ctx context.Context) (*ProfileCache, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // use default Addr
		Password: "",               // no password set
		DB:       0,                // use default DB
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return &ProfileCache{
		DB: client,
	}, err
}

func (c *ProfileCache) Close() {
	c.DB.Close()
}
