package setup

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// DB setup the database connection and do checks
func DB(dbname string) *sql.DB {
	db := setupDb(dbname)
	logSuccess(dbname)
	return db
}

func setupDb(dbname string) *sql.DB {
	checkFile(dbname)
	db, err := sql.Open("sqlite3", fmt.Sprintf("./db/%s", dbname))
	if err != nil {
		log.Default().Fatalln("Cannot open the DB. On DB Func ->", err.Error())
	}
	if err = db.Ping(); err != nil {
		log.Default().Fatalln("Cannot ping the DB, maybe is offline. On DB Func ->", err.Error())
	}
	checkTableState(db, "DISCORD_USERS")
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	return db
}

func checkFile(dbname string) {
	fullpath := fmt.Sprintf("./db/%s", dbname)
	if _, err := os.Stat(fullpath); os.IsNotExist(err) {
		log.Default().Fatalln("DB file does not exist:", err)
	}
	// Check file permissions
	if info, err := os.Stat(fullpath); err == nil {
		mode := info.Mode()
		expectedMode := os.FileMode(0666)
		if mode != expectedMode {
			log.Fatalln(fmt.Sprintf("Error: File %s has incorrect permissions %s", fullpath, mode))
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

func logSuccess(dbname string) {
	log.Default().Printf("DB Current tables state is ok.")
	log.Default().Printf("DB Current permission state is ok.")
	log.Default().Printf("Connect estabilished with DB: ./db/%s", dbname)
}
