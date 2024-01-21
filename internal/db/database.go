package db

import (
	"database/sql"
	"log"
	"time"

	"github.com/leoff00/ta-pago-bot/pkg/env"
	_ "github.com/mattn/go-sqlite3"
)

func Setup() {
	db, err := sql.Open("sqlite3", "ta_pago.db")

	if err != nil {
		log.Default().Fatalln("Cannot open the DB. On Setup Func ->", err.Error())
	}

	if err = db.Ping(); err != nil {
		log.Default().Fatalln("Cannot ping the DB, maybe is offline. On Setup Func ->", err.Error())
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	log.Default().Printf("Connect estabilished with DB: %s On Setup DB", env.Getenv("DB_NAME"))

	defer db.Close()
}
