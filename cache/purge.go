package cache

import (
	"fmt"

	"github.com/Noah-Huppert/gh-gantt/zenhub"

	"github.com/go-redis/cache"
	"github.com/go-redis/redis"
)

// PurgeCache deletes the specified cache's contents.
//
// An error is returned if one occurs.
func PurgeCache(redisClient *redis.Client, redisCache *cache.Codec,
	cacheName string) error {

	// Check if cacheName is indicating a category of models stored in the
	// cache and not a specific redis key to clear
	if cacheName == zenhub.DepsCategoryKey {
		err := PurgeCacheCategory(redisClient, redisCache, zenhub.DepKeysKey)
		if err != nil {
			return fmt.Errorf("error deleting cache category: %s",
				err.Error())
		}
	}

	// Otherwise clear cache normally
	if err := redisCache.Delete(cacheName); err != nil {
		return fmt.Errorf("error deleting cache: %s, err: %s",
			cacheName, err.Error())
	}

	return nil
}

// PurgeCacheCategory purges all instances of a model from the cache. The name
// of a Redis set holding the keys of all the model instances should be
// provided.
//
// An error is returned if one occurs.
func PurgeCacheCategory(redisClient *redis.Client, redisCache *cache.Codec,
	keysKey string) error {

	// Get members of instance keys set
	keys, err := redisClient.SMembers(keysKey).Result()
	if err != nil {
		return fmt.Errorf("error retrieving instance keys from set: "+
			"%s, err: %s", keysKey, err.Error())
	}

	// Delete each key
	for _, key := range keys {
		err := PurgeCache(redisClient, redisCache, key)
		if err != nil {
			return fmt.Errorf("error deleting cache category key: "+
				"%s, err: %s", key, err.Error())
		}
	}

	// Success
	return nil
}
