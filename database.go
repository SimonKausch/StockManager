package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Stock struct {
	ID              int64
	XLength         int
	YLength         int
	ZLength         int
	Material        string
	CertificatePath string
	InvoicePath     string
	Location        string
}

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
			location TEXT,
			certificate_path TEXT,
			invoice_path TEXT
		)
	`)
	if err != nil {
		return nil, fmt.Errorf("error creating table: %v", err)
	}

	return db, nil
}

func insertStock(stock *Stock) error {
	query := `
		INSERT INTO stock (xLength, yLength, zLength, material, certificate_path, invoice_path)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	_, err := Db.Exec(query, stock.XLength, stock.YLength, stock.ZLength, stock.Material, stock.CertificatePath, stock.InvoicePath)
	return err
}

func ListStock() ([]Stock, error) {
	row, err := Db.Query("SELECT * FROM stock ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer row.Close()

	stockMaterials := []Stock{}
	for row.Next() {
		var stock Stock
		err := row.Scan(&stock.ID, &stock.XLength, &stock.YLength, &stock.ZLength, &stock.Material, &stock.CertificatePath, &stock.InvoicePath)
		if err != nil {
			return nil, err
		}
		stockMaterials = append(stockMaterials, stock)
	}
	return stockMaterials, nil
}

func addStock(stock Stock) {
	err := insertStock(&stock)
	if err != nil {
		log.Fatal(err)
	}
}
