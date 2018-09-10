package test

import (
	"github.com/Noah-Huppert/gh-gantt/server/config"
	"github.com/Noah-Huppert/gh-gantt/server/libdb"

	"github.com/jmoiron/sqlx"
)

// ConnTestDB connects to the database
func ConnTestDB() (*sqlx.DB, error) {
	dbCfg := config.DBConfig{
		DBHost:     "localhost",
		DBPort:     5432,
		DBName:     "test-gh-gantt",
		DBUsername: "test-gh-gantt",
	}

	return libdb.ConnectX(dbCfg)
}
