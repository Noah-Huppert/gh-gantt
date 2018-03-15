package server

import (
	"fmt"
	"net/http"

	"github.com/Noah-Huppert/gh-gantt/cache"
	"github.com/Noah-Huppert/gh-gantt/gh"
	"github.com/Noah-Huppert/gh-gantt/zenhub"

	cacheLib "github.com/go-redis/cache"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

// PurgePath is the path to register Purge handlers at
const PurgePath string = "/api/cache/purge"

// PurgeEndpoint implements HTTP handlers for the purge endpoint
type PurgeEndpoint struct {
	// BasePath is the URL HTTP Purge handlers will be registered at
	BasePath string

	// redisClient is the Redis client used to purge parts of the redis cache
	redisClient *redis.Client

	// redisCache is the Redis client used to access the cache store
	redisCache *cacheLib.Codec
}

// NewPurgeEndpoint creates a new PurgeEndpoint instance
func NewPurgeEndpoint(redisClient *redis.Client, redisCache *cacheLib.Codec) *PurgeEndpoint {

	return &PurgeEndpoint{
		BasePath:    PurgePath,
		redisClient: redisClient,
		redisCache:  redisCache,
	}
}

// Register implements Registerable.Register
func (p PurgeEndpoint) Register(router *mux.Router) {
	router.HandleFunc(p.BasePath, p.Post).Methods("POST")
}

// allowedCacheNames are the values allowed in the `caches` body parameter
var allowedCacheNames []string = []string{gh.IssuesCacheKey, gh.RepoCacheKey,
	zenhub.DepsCategoryKey}

// purgeBody is the structure purge post request bodies will be serialized into.
type purgeBody struct {
	// Caches holds the names of the caches to purge
	Caches []string `json:"caches"`
}

// Post handles purge endpoint post requests
func (p PurgeEndpoint) Post(w http.ResponseWriter, r *http.Request) {
	if RequireBodyFields(w, r, []string{"caches"}) {
		// Parse body
		body := purgeBody{}
		err := ParseBody(r, &body)

		if err != nil {
			WriteErr(w, 500, fmt.Errorf("error reading body: %s",
				err.Error()))
			return
		}

		// Delete
		purged := []string{}

		for _, cacheName := range body.Caches {
			// Ensure valid cache name
			valid := false
			for _, validName := range allowedCacheNames {
				if validName == cacheName {
					valid = true
				}
			}

			if !valid {
				WriteErr(w, 500, fmt.Errorf("unknown cache: %s",
					cacheName))
				return
			}

			if err := cache.PurgeCache(p.redisClient, p.redisCache,
				cacheName); err != nil {

				WriteErr(w, 500, fmt.Errorf("error "+
					"purging cache: %s, err: %s",
					cacheName, err.Error()))
				return
			}

			purged = append(purged, cacheName)
		}

		// Success
		resp := map[string]interface{}{
			"purged": purged,
			"errors": []string{},
		}
		WriteJSON(w, 200, resp)
	}
}
