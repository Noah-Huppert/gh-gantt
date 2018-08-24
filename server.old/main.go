package main

import (
	"fmt"
	"net/http"

	"github.com/Noah-Huppert/gh-gantt/server/api"
	"github.com/Noah-Huppert/gh-gantt/server/config"

	"github.com/Noah-Huppert/golog"
	"github.com/ekyoung/gin-nice-recovery"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	logger := golog.NewStdLogger("gh-gantt")

	// Load configuration
	cfg, err := config.NewConfig()
	if err != nil {
		logger.Fatalf("error loading configuration: %s", err.Error())
	}

	// Setup routes
	router := gin.Default()

	router.Use(nice.Recovery(api.RecoveryHandler))

	v0API := router.Group("/api/v0")

	router.Use(static.Serve("/", static.LocalFile("../frontend/dist", true)))

	v0API.GET("/healthz", api.HealthCheckHandler)
	v0API.GET("/auth/login", api.MakeAuthLoginHandler(cfg.GithubClientID))
	v0API.POST("/auth/callback", api.MakeAuthCallbackHandler(cfg.GithubClientID, cfg.GithubClientSecret))

	err = endless.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), router)
	if err != nil && err != http.ErrServerClosed {
		logger.Fatalf("error running HTTP server: %s", err.Error())
	}
}
