package database

import (
	"database/sql"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

func DbConnection() *sql.DB {

	basePath, _ := os.Getwd()
	dbPath := filepath.Join(basePath, "../../internal/database/tasks.db")

	DB, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil
	}

	return DB
}
