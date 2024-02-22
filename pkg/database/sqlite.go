package database

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"time"
)

const (
	DB_FOLDER         = "/db"
	MIGRATIONS_FOLDER = "/db/migrations"
)

// SetupSqlite do connection, run migrations and do state checks
func SetupSqlite(dbname string, tablesToCheck []string, env string) *sql.DB {
	dbPath := "." + DB_FOLDER + "/" + dbname
	doMigrations(dbname)

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Default().Fatalf(fmt.Sprintf("DB: Cannot open %s. Err ->", dbname), err.Error())
	}
	if err = db.Ping(); err != nil {
		log.Default().Fatalf(fmt.Sprintf("DB: Cannot ping on %s. Err ->", dbname), err.Error())
	}

	checkFile(dbPath, dbname, env) // To avoid permission issues on prod env with sqlite, docker and volumes

	for _, table := range tablesToCheck {
		checkTableState(db, table, dbname)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	logSuccess(dbname)
	return db
}

func checkFile(path, dbname, env string) {
	if env != "PROD" {
		log.Default().Printf("DB: Skipping file permission check on %s with non-PROD env ", dbname)
		return
	}
	info, err := os.Stat(path)
	if err != nil {
		log.Fatalf("DB: Error accessing file %s: %s", dbname, err)
	}
	mode := info.Mode()
	expectedMode := os.FileMode(0666)
	if mode != expectedMode {
		if err := os.Chmod(path, expectedMode); err != nil {
			log.Fatalf(`Error: File %s has incorrect permissions %s and can't change it to %s automatically`, path, mode, expectedMode)
		}
		log.Printf("DB: File %s permissions changed to %s", path, expectedMode)
	}
	log.Default().Printf("DB: Current %s permission state is ok.", dbname)
}

func checkTableState(db *sql.DB, tableName string, dbname string) {
	query := fmt.Sprintf("SELECT name FROM sqlite_master WHERE type='table' AND name='%s'", tableName)
	rows, err := db.Query(query)
	if err != nil {
		log.Default().Fatalln("Cannot check if the table exists. On DB Func ->", err.Error())
	}
	defer rows.Close()
	tableExists := rows.Next()
	if !tableExists {
		log.Default().Fatalf("DB: Table %s does not exist on %s", tableName, dbname)
	}
}

func doMigrations(dbname string) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get current working directory: %s", err.Error())
	}
	migrationsPath := cwd + MIGRATIONS_FOLDER
	dbPath := cwd + DB_FOLDER + "/" + dbname
	m, err := migrate.New(
		"file://"+migrationsPath,
		"sqlite://"+dbPath,
	)
	if err != nil {
		log.Default().Fatalf("failed to create migration instance: %v", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Default().Fatalf("DB: failed to apply migrations for %s: %v on path", dbname, err)
	}
}

func logSuccess(dbname string) {
	log.Default().Printf("DB: Current %s tables state is ok.", dbname)
	log.Default().Printf("DB: Connect estabilished with: %s", dbname)
}
