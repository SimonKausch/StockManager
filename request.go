package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	serverPort   = "8000"
	serverAdress = "http://127.0.0.1"
	endpoint     = "/box/model.stp"
)

func main() {
	res, err := http.Get(serverAdress + ":" + serverPort + endpoint)
	if err != nil {
		log.Println(err)
		os.Exit(0)
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		os.Exit(0)
	}
	if string(body) == "null" {
		log.Println("File not found")
		os.Exit(0)
	}

	var numbers []float64
	err = json.Unmarshal(body, &numbers)

	res.Body.Close()

	fmt.Println(numbers)
}
