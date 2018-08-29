package main

import (
	"database/sql"
	"fmt"

	"github.com/Noah-Huppert/gh-gantt/server/config"

	"github.com/Noah-Huppert/golog"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
)

func main() {
	// Setup logger
	logger := golog.NewStdLogger("db-migrate")

	// Load configuration
	cfg, err := config.NewDBConfig()
	if err != nil {
		logger.Fatalf("failed to load database configuration: %s", err.Error())
	}

	// Connect to database
	dbConnStr := fmt.Sprintf("postgres://%s", cfg.DBUsername)

	if len(cfg.DBPassword) > 0 {
		dbConnStr += fmt.Sprintf(":%s", cfg.DBPassword)
	}

	dbConnStr += fmt.Sprintf("@%s:%d/%s?sslmode=disable", cfg.DBHost, cfg.DBPort, cfg.DBName)

	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		logger.Fatalf("failed to connect to database: %s", err.Error())
	}

	// Make database driver
	dbDriver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		logger.Fatalf("error creating database driver: %s", err.Error())
	}

	// Create migrate client
	migrator, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", dbDriver)
	if err != nil {
		logger.Fatalf("error creating migrator: %s", err.Error())
	}

	// Migrate
	err = migrator.Up()
	if err != nil {
		logger.Fatalf("error running migrations: %s", err.Error())
	}

	logger.Info("success")
}
