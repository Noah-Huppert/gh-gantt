package server

import (
	"fmt"
	"net/http"

	"github.com/Noah-Huppert/gh-gantt/gh"

	"github.com/go-redis/cache"
	"github.com/gorilla/mux"
)

// PurgePath is the path to register Purge handlers at
const PurgePath string = "/api/cache/purge"

// PurgeEndpoint implements HTTP handlers for the purge endpoint
type PurgeEndpoint struct {
	// BasePath is the URL HTTP Purge handlers will be registered at
	BasePath string

	// redisCache is the Redis client used to access the cache store
	redisCache *cache.Codec
}

// NewPurgeEndpoint creates a new PurgeEndpoint instance
func NewPurgeEndpoint(redisCache *cache.Codec) *PurgeEndpoint {
	return &PurgeEndpoint{
		BasePath:   PurgePath,
		redisCache: redisCache,
	}
}

// Register implements Registerable.Register
func (p PurgeEndpoint) Register(router *mux.Router) {
	router.HandleFunc(p.BasePath, p.Post).Methods("POST")
}

// zenhubDepsCache is the value passed in the `caches` body field to clear all
// ZenHub dependencies from the cache
const zenhubDepsCache string = "zenhub.dependencies"

// allowedCacheNames are the values allowed in the `caches` body parameter
var allowedCacheNames []string = []string{gh.IssuesCacheKey, gh.RepoCacheKey,
	zenhubDepsCache}

// Post handles purge endpoint post requests
func (p PurgeEndpoint) Post(w http.ResponseWriter, r *http.Request) {
	if RequireBodyFields(w, r, []string{"caches"}) {
		caches := r.Form["caches"]

		// Delete
		fmt.Fprintf(w, "purge %s", caches)
	}
}
