package main

import (
	"bytes"
	"fmt"
	"image/color"
	"strconv"
	"text/tabwriter"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

// TODO: Import materials from file
var listMaterials []string = []string{"1.2313", "1.4305", "1.4404", "Aluminium", "Cu"}

// createTable creates an output in table form from
// a slice of Stock
func createTable(stock []Stock) string {
	var buf bytes.Buffer

	w := tabwriter.NewWriter(&buf, 10, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tX\tY\tZ\tMaterial")
	for _, s := range stock {
		out := strconv.Itoa(int(s.ID)) + "\t"
		out += strconv.Itoa(int(s.XLength)) + "\t"
		out += strconv.Itoa(int(s.YLength)) + "\t"
		out += strconv.Itoa(int(s.ZLength)) + "\t"
		out += s.Material
		fmt.Fprintln(w, out)
	}
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

func CreateTitleText(s string) *canvas.Text {
	t := canvas.NewText(s, color.Black)
	t.TextStyle = fyne.TextStyle{Bold: true}
	t.Alignment = fyne.TextAlignCenter
	t.TextSize = 20

	return t
}
