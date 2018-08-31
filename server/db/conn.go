package db

import (
	"fmt"

	"github.com/Noah-Huppert/gh-gantt/server/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect(dbCfg config.DBConfig) (*sqlx.DB, error) {
	sqlConnStr := fmt.Sprintf("host=%s port=%d dbname=%s user=%s sslmode=disable", dbCfg.DBHost, dbCfg.DBPort,
		dbCfg.DBName, dbCfg.DBUsername)

	if len(dbCfg.DBPassword) > 0 {
		sqlConnStr += fmt.Sprintf("password=%s", dbCfg.DBPassword)
	}

	db, err := sqlx.Connect("postgres", sqlConnStr)
	if err != nil {
		return nil, fmt.Errorf("error connecting to the database: %s", err.Error())
	}

	return db, nil
}
