package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"text/tabwriter"
)

func createTable(stock []Stock) string {
	var buf bytes.Buffer

	w := tabwriter.NewWriter(&buf, 10, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tX\tY\tZ\tMaterial")
	// TODO: for loop over slice of Stock
	for _, s := range stock {
		out := strconv.Itoa(int(s.ID)) + "\t"
		out += strconv.Itoa(int(s.XLength)) + "\t"
		out += strconv.Itoa(int(s.YLength)) + "\t"
		out += strconv.Itoa(int(s.ZLength)) + "\t"
		out += s.Material
		fmt.Fprintln(w, out)
	}
	// fmt.Fprintln(w, "10\t100\t50\t30\tStahl")
	w.Flush()

	return buf.String()
}

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
	t += "\tX: "
	t += strconv.Itoa(s.XLength)
	t += "\tY: "
	t += strconv.Itoa(s.YLength)
	t += "    Z: "
	t += strconv.Itoa(s.ZLength)
	t += "    Material: "
	t += s.Material
	t += "    Location: "
	t += s.Location
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
