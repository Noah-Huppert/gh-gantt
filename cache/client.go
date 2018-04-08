package cache

import (
	"fmt"
	"time"

	"github.com/Noah-Huppert/gh-gantt/config"

	"github.com/go-redis/cache"
	redisLib "github.com/go-redis/redis"
	"github.com/vmihailenco/msgpack"
)

// marshal encodes an object for storage in redis
func marshal(v interface{}) ([]byte, error) {
	return msgpack.Marshal(v)
}

// unmarshal decodes an object from redis storage
func unmarshal(b []byte, v interface{}) error {
	return msgpack.Unmarshal(b, v)
}

// instance caches a cache.Codec client
var instance *cache.Codec

// pingKey is the Redis key the Redis client will attempt to set to verify it
// it is connected correctly.
const pingKey string = "internal.ping"

// NewClient creates a new Redis cache codec
func NewClient(cfg *config.Config) (*cache.Codec, error) {
	// If already created
	if instance != nil {
		return instance, nil
	}

	ring := redisLib.NewRing(&redisLib.RingOptions{
		Addrs: map[string]string{
			"server": cfg.Redis.Host,
		},
	})

	codec := &cache.Codec{
		Redis: ring,

		Marshal:   marshal,
		Unmarshal: unmarshal,
	}

	// Test connection
	dur, err := time.ParseDuration("1s")
	if err != nil {
		return nil, fmt.Errorf("error creating Redis connection test"+
			" expiration duration: %s", err.Error())
	}

	err = codec.Set(&cache.Item{
		Key:        pingKey,
		Object:     "ping",
		Expiration: dur,
	})
	if err != nil {
		return nil, fmt.Errorf("error connecting to Redis host: %s",
			err.Error())
	}

	instance = codec
	return instance, nil
}
