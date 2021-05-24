package redis

import (
	"context"
	"os"

	goredis "github.com/go-redis/redis/v8"
)

type Redis struct {
	client goredis.UniversalClient
}

func NewRedis() Redis {
	universalClient := goredis.NewClient(&goredis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
	})

	return Redis{client: universalClient}
}

func (r Redis) Get(ctx context.Context, key string) (string, error) {
	cacheValue, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return cacheValue, nil
}

func (r Redis) Set(ctx context.Context, key, value string) error {
	err := r.client.Set(ctx, key, value, 0).Err()
	if err != nil {
		return err
	}

	return nil
}
