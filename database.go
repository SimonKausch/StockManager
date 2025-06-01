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

func addStock(stock *Stock) error {
	query := `
		INSERT INTO stock (xLength, yLength, zLength, material, certificate_path, invoice_path)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	_, err := Db.Exec(query, stock.XLength, stock.YLength, stock.ZLength, stock.Material, stock.CertificatePath, stock.InvoicePath)
	return err
}

func searchStock(stock *Stock) ([]Stock, error) {
	// TODO: Check for the correct material
	// TODO: Add test for this function

	var result []Stock

	query := `
        SELECT ID, xLength, yLength, zLength, material
        FROM stock
        WHERE xLength >= ? AND yLength >= ? AND zLength >= ?
    `
	args := []any{stock.XLength, stock.YLength, stock.ZLength}

	// Check material if not empty
	if stock.Material != "" {
		query += " AND material = ?"
		args = append(args, stock.Material)
	}

	rows, err := Db.Query(query, args...)
	if err != nil {
		log.Printf("Error querying stock by length: %v", err)
		return nil, err
	}

	for rows.Next() {
		var s Stock
		// Scan the columns into the fields of the Stock struct
		// Make sure the order and types match your SELECT statement
		err := rows.Scan(&s.ID, &s.XLength, &s.YLength, &s.ZLength, &s.Material)
		if err != nil {
			log.Printf("Error scanning stock row: %v", err)
			return nil, err // Or handle the error as appropriate
		}
		result = append(result, s)
		log.Printf("Found suitable material: %v", s)
	}

	// Check for errors encountered during iteration
	if err = rows.Err(); err != nil {
		log.Printf("Error during rows iteration: %v", err)
		return nil, err
	}

	return result, nil
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
