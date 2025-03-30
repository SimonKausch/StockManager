package main

func main() {
	db, err := initDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	for {
		gui()
	}

	// // List all stock in db
	// addStock(db)
	// allStock, err := ListStock(db)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// for _, stock := range allStock {
	// 	fmt.Println(stock)
	// }
}
