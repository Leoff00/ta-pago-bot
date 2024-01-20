package db

import (
	"database/sql"
	"log"
	"time"

	"github.com/leoff00/ta-pago-bot/pkg/env"
	_ "github.com/mattn/go-sqlite3"
)

func Setup() {
	database, err := sql.Open("sqlite3", "ta_pago.db")

	if err != nil {
		log.Default().Fatalln("Cannot open the DB. On Setup Func ->", err.Error())
	}

	if err = database.Ping(); err != nil {
		log.Default().Fatalln("Cannot ping the DB, maybe is offline. On Setup Func ->", err.Error())
	}

	database.SetMaxIdleConns(10)
	database.SetMaxOpenConns(100)
	database.SetConnMaxLifetime(time.Hour)

	log.Default().Printf("Connect estabilished with DB: %s On Setup DB", env.Getenv("DB_NAME"))

}
