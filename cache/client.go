package cache

import (
	"github.com/Noah-Huppert/gh-gantt/config"

	redisLib "github.com/go-redis/redis"
	"github.com/vmihailenco/msgpack"

	"github.com/go-redis/cache"
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

// NewClient creates a new Redis cache codec
func NewClient(cfg *config.Config) *cache.Codec {
	// If already created
	if instance != nil {
		return nil
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

	instance = codec
	return instance
}
