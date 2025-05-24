package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// parseIntInput attempts to convert text to an integer and returns a
// more informative error if it fails.
func parseIntInput(text string) (int, error) {
	val, err := strconv.Atoi(text)
	if err != nil || val <= 0 {
		// Return a new error that includes the name of the input field
		return 0, fmt.Errorf("invalid integer for %s: %w", text, err)
	}
	return val, nil
}

func getFilesinDir() ([]string, error) {
	entries, err := os.ReadDir(stepDir)
	if err != nil {
		return nil, err
	}

	var files []string

	for _, entry := range entries {
		files = append(files, entry.Name())
	}
	return files, nil
}

func PrintStock(s Stock) string {
	var t string
	t += strconv.Itoa(int(s.ID))
	t += "    X: "
	t += strconv.Itoa(s.XLength)
	t += "    Y: "
	t += strconv.Itoa(s.YLength)
	t += "    Z: "
	t += strconv.Itoa(s.ZLength)
	t += "    Material: "
	t += s.Material
	t += "\n"

	return t
}

// Return stock as one string
func listStockString() *string {
	var out string

	slice, err := ListStock()
	if err != nil {
		log.Fatal(err)
	}
	for _, s := range slice {
		out += PrintStock(s)
	}
	return &out
}

// Return stock as buttons
func listStockButtons() []fyne.CanvasObject {
	// TODO: Use this in gui.go
	var items []fyne.CanvasObject

	slice, err := ListStock()
	if err != nil {
		log.Fatal(err)
	}

	for i, s := range slice {
		label := strconv.Itoa(int(s.ID)) + " " + s.Material
		items = append(items, widget.NewButton(label, func() {
			fmt.Println("Tapped", i)
		}))
	}

	return items
}
