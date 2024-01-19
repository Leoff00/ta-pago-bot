package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/leoff00/ta-pago-bot/pkg/env"
	_ "github.com/lib/pq"
)

const dbDriver = "postgres"

func Setup() (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"%s://%s:%s@%s/%s?sslmode=disable", dbDriver,
		env.Getenv("DB_USER"),
		env.Getenv("DB_PASSWORD"),
		env.Getenv("DB_HOST"),
		env.Getenv("DB_NAME"),
	)
	db, err := sql.Open(dbDriver, connStr)
	log.Printf("Database '%v' connection established", env.Getenv("DB_NAME"))
	if err != nil {
		return nil, err
	}
	return db, nil
}
