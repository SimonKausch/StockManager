package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func initDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "stock.db")
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	// Create tables if they don't exist
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS stock (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			xLength INTEGER NOT NULL,
			yLength INTEGER NOT NULL,
			zLength INTEGER NOT NULL,
			material TEXT NOT NULL,
			certificate_path TEXT,
			invoice_path TEXT
		)
	`)
	if err != nil {
		return nil, fmt.Errorf("error creating table: %v", err)
	}

	return db, nil
}
