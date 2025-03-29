package main

import (
	"database/sql"
	"fmt"
	"log"
)

func main() {
	db, err := initDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Your application logic here
	addStock(db)

	allStock, err := ListStock(db)
	if err != nil {
		log.Fatal(err)
	}

	for _, stock := range allStock {
		fmt.Println(stock)
	}
}

func addStock(db *sql.DB) {
	stock := Stock{
		XLength:         50,
		YLength:         75,
		ZLength:         20,
		Material:        "1.4301",
		CertificatePath: "id2202.pdf",
		InvoicePath:     "invoice.pdf",
	}
	err := insertStock(db, &stock)
	if err != nil {
		log.Fatal(err)
	}
}
