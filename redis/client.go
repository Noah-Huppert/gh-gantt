package redis

import (
	"fmt"

	"github.com/Noah-Huppert/gh-gantt/config"

	"github.com/go-redis/redis"
)

// instance caches a redis.Client
var instance *redis.Client

// NewClient returns a new Redis client
func NewClient(cfg *config.Config) (*redis.Client, error) {
	// If already created
	if instance != nil {
		return instance, nil
	}

	// Create client
	client := redis.NewClient(&redis.Options{
		Addr: cfg.Redis.Host,
	})

	// Test connection
	_, err := client.Ping().Result()
	if err != nil {
		return nil, fmt.Errorf("error connecting to Redis host: %s",
			err.Error())
	}

	instance = client

	// Success
	return instance, nil
}
