package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
)

const (
	serverPort   = "8000"
	serverAdress = "http://127.0.0.1"
	endpoint     = "/box/model.stp"
)

type boundingBox struct {
	boxX float64
	boxY float64
	boxZ float64
}

func bboxLengths(n []float64) boundingBox {
	var b boundingBox

	b.boxX = n[3] - n[0]
	b.boxY = n[4] - n[1]
	b.boxZ = n[5] - n[2]

	b.boxX = math.Ceil(b.boxX*100) / 100
	b.boxY = math.Ceil(b.boxY*100) / 100
	b.boxZ = math.Ceil(b.boxZ*100) / 100

	return b
}

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

	log.Println(numbers)
	log.Println(bboxLengths(numbers))
}
