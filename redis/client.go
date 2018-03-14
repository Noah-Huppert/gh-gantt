package redis

import (
	"github.com/Noah-Huppert/gh-gantt/config"

	"github.com/go-redis/redis"
)

// NewClient returns a new Redis client
func NewClient(cfg *config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: cfg.Redis.Host,
	})
}
