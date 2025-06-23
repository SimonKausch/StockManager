package main

import "database/sql"

var DB *sql.DB

func main() {
	var err error

	DB, err = initDB()
	if err != nil {
		panic(err)
	}
	defer DB.Close()

	for {
		gui()
	}
}
