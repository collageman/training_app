// auth-service/pkg/config/redis.go
package config

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func SetupRedis() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", Config.Redis.Host, Config.Redis.Port),
		Password: Config.Redis.Password,
		DB:       Config.Redis.DB,
	})

	// Test connection
	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to redis: %w", err)
	}

	RedisClient = client
	return client, nil
}
