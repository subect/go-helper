package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

func InitializeRedisClient(address string, password string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
	})
	return client
}

func SetKey(client *redis.Client, ctx context.Context, key string, value string) error {
	err := client.Set(ctx, key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetKey(client *redis.Client, ctx context.Context, key string) (string, error) {
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
