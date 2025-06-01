package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"math"
	"net/http"
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
func requestBBox(filename string) (BoundingBox, error) {
	endpoint := "/box/" + filename

	var bBox BoundingBox

	res, err := http.Get(serverAdress + ":" + serverPort + endpoint)
	if err != nil {
		return bBox, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return bBox, err
	}
	if string(body) == "null" {
		err = errors.New("file not found")
		return bBox, err
	}

	var numbers []float64
	err = json.Unmarshal(body, &numbers)
	if err != nil {
		return bBox, err
	}

	res.Body.Close()

	log.Println(numbers)

	bBox = bboxLengths(numbers)

	return bBox, nil
}
