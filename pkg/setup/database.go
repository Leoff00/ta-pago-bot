package setup

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/leoff00/ta-pago-bot/pkg/env"
	_ "github.com/mattn/go-sqlite3"
)

var (
	DB_PATH = env.Getenv("DB_PATH")
	DB_NAME = env.Getenv("DB_NAME")
)

func DB() *sql.DB {
	db := setupDb(DB_PATH, DB_NAME, "DISCORD_USERS")
	logSuccess()
	return db
}

func setupDb(dbpath string, dbname string, tableName string) *sql.DB {
	checkFile(dbpath, dbname)
	db, err := sql.Open("sqlite3", fmt.Sprintf("%s/%s", dbpath, dbname))
	if err != nil {
		log.Default().Fatalln("Cannot open the DB. On DB Func ->", err.Error())
	}
	if err = db.Ping(); err != nil {
		log.Default().Fatalln("Cannot ping the DB, maybe is offline. On DB Func ->", err.Error())
	}
	checkTableState(db, tableName)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour) // idk what happens when it expires
	return db
}

func checkFile(path string, dbname string) {
	fullpath := fmt.Sprintf("%s/%s", path, dbname)
	if _, err := os.Stat(fullpath); os.IsNotExist(err) {
		log.Fatalln("DB file does not exist:", err)
	}
	// Check file permissions
	if info, err := os.Stat(fullpath); err == nil {
		mode := info.Mode()
		// Check if the file permissions are not equal to -rw-rw-rw-
		expectedMode := os.FileMode(0666)
		if mode != expectedMode {
			log.Printf("File %s has incorrect permissions: %s", fullpath, mode)
			log.Fatalln(fmt.Sprintf("Error: File %s has incorrect permissions", fullpath))
		}
	}
}

func checkTableState(db *sql.DB, tableName string) {
	query := fmt.Sprintf("SELECT name FROM sqlite_master WHERE type='table' AND name='%s'", tableName)
	rows, err := db.Query(query)
	if err != nil {
		log.Default().Fatalln("Cannot check if the table exists. On DB Func ->", err.Error())
	}
	defer rows.Close()
	tableExists := rows.Next()
	if !tableExists {
		log.Default().Fatalln("Table DISCORD_USERS does not exist")
	}
}

func logSuccess() {
	log.Default().Printf("DB Current tables state is ok.")
	log.Default().Printf("DB Current permission state is ok.")
	log.Default().Printf("Connect estabilished with DB: %s/%s", DB_PATH, DB_NAME)
}
