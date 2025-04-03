package main

import "database/sql"

var Db *sql.DB

func main() {
	var err error

	Db, err = initDB()
	if err != nil {
		panic(err)
	}
	defer Db.Close()

	for {
		gui()
	}
}
