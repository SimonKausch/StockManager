package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"math"
	"mime/multipart"
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

func uploadFile(filepath string) (BoundingBox, error) {
	var bBox BoundingBox

	// Open the file
	file, err := os.Open(filepath)
	if err != nil {
		return bBox, err
	}
	defer file.Close()

	// Create a new multipart writer
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Create the form file field
	part, err := writer.CreateFormFile("file", filepath)
	if err != nil {
		return bBox, err
	}

	// Copy the file content to the form field
	_, err = io.Copy(part, file)
	if err != nil {
		return bBox, err
	}

	// Close the multipart writer to finalize the body
	err = writer.Close()
	if err != nil {
		return bBox, err
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", "http://localhost:8000/upload/", body)
	if err != nil {
		return bBox, err
	}

	// Set the Content-Type header with the boundary
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Send the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return bBox, err
	}
	defer res.Body.Close()

	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		return bBox, err
	}
	if string(responseBody) == "null" {
		err = errors.New("file not found")
		return bBox, err
	}

	var numbers []float64
	err = json.Unmarshal(responseBody, &numbers)
	if err != nil {
		return bBox, err
	}

	res.Body.Close()

	log.Println(numbers)

	bBox = bboxLengths(numbers)

	return bBox, nil
}
