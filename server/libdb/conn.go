package libdb

import (
	"database/sql"
	"fmt"

	"github.com/Noah-Huppert/gh-gantt/server/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Connect to a PostgreSQL database
func Connect(dbCfg config.DBConfig) (*sql.DB, error) {
	sqlConnStr := fmt.Sprintf("host=%s port=%d dbname=%s user=%s sslmode=disable", dbCfg.DBHost, dbCfg.DBPort,
		dbCfg.DBName, dbCfg.DBUsername)

	if len(dbCfg.DBPassword) > 0 {
		sqlConnStr += fmt.Sprintf("password=%s", dbCfg.DBPassword)
	}

	db, err := sql.Open("postgres", sqlConnStr)
	if err != nil {
		return nil, fmt.Errorf("error connecting to the database: %s", err.Error())
	}

	return db, nil
}
