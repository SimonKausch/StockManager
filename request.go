package main

import (
	"encoding/json"
	"io"
	"log"
	"math"
	"net/http"
	"os"
)

const (
	serverPort   = "8000"
	serverAdress = "http://127.0.0.1"
)

type BoundingBox struct {
	BoxX float64
	BoxY float64
	BoxZ float64
}

// Calculates the lengts of the bounding box in x,y,z
// Returns type BoundingBox
func bboxLengths(n []float64) BoundingBox {
	var b BoundingBox

	b.BoxX = n[3] - n[0]
	b.BoxY = n[4] - n[1]
	b.BoxZ = n[5] - n[2]

	b.BoxX = math.Ceil(b.BoxX*100) / 100
	b.BoxY = math.Ceil(b.BoxY*100) / 100
	b.BoxZ = math.Ceil(b.BoxZ*100) / 100

	return b
}

// Requests the bounding box in x, y and z axis
// Returns BoundingBox type
func requestBBox(filename string) BoundingBox {
	endpoint := "/box/" + filename

	res, err := http.Get(serverAdress + ":" + serverPort + endpoint)
	if err != nil {
		log.Println(err)
		// TODO: Error message instead of os.Exit
		os.Exit(0)
	}

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
	if err != nil {
		log.Println(err)
	}

	res.Body.Close()

	log.Println(numbers)

	bBox := bboxLengths(numbers)

	return bBox
}
