package main

import (
	"fmt"
	"image/color"
	"log"
	"os"
	"strconv"
	"time"

	"fyne.io/fyne/v2/layout"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const stepDir = "../stepReader/files/"

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

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)
}

func printStock(s Stock) string {
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

func listStockGUI() *string {
	var out string

	slice, err := ListStock()
	if err != nil {
		log.Fatal(err)
	}
	for _, s := range slice {
		out += printStock(s)
	}
	return &out
}

func gui() {
	a := app.New()
	w := a.NewWindow("StockManger")
	w.Resize(fyne.Size{Width: 750, Height: 500})

	linesep := canvas.NewLine(color.Black)
	allStock := widget.NewLabel("")
	listStockButton := widget.NewButton("List all stock", func() {
		allStock.SetText(*listStockGUI())
	})
	addStockButton := widget.NewButton("Add or search Stock", func() {
		addStockWindow(a)
	})
	bboxStepButton := widget.NewButton("Analyze step file", func() {
		addStepWindow(a)
	})

	w.SetContent(container.NewVBox(linesep, allStock, listStockButton, addStockButton, bboxStepButton))

	w.ShowAndRun()
}

func addStepWindow(a fyne.App) {
	wAdd := a.NewWindow("Get bounding box")
	wAdd.Resize(fyne.NewSize(400, 400))

	listFiles, err := getFilesinDir()
	if err != nil {
		output := widget.NewLabel(fmt.Sprint(err))
		wAdd.SetContent(container.NewVBox(output))
		wAdd.Show()
	} else {
		// var r string
		output := widget.NewRadioGroup(listFiles, func(r string) {
		})
		analyzeFile := widget.NewButton("Analyze chosen file", func() {
			// TODO: replace filename with result from NewRadioGroup
			bbox := requestBBox("model.stp")

			// From float to a string with a precision of 1 decimal place
			x := fmt.Sprintf("X: %.1f\n", bbox.BoxX)
			y := fmt.Sprintf("X: %.1f\n", bbox.BoxY)
			z := fmt.Sprintf("X: %.1f\n", bbox.BoxZ)

			result := x + y + z
			outputBox := widget.NewLabel(result)

			// Update window
			wAdd.SetContent(outputBox)
			wAdd.Show()
		})
		wAdd.SetContent(container.NewVBox(output, analyzeFile))
		wAdd.Show()
	}
}

func addStockWindow(a fyne.App) {
	// TODO: Show list of materials from json file
	wAdd := a.NewWindow("Add stock material")
	wAdd.Resize(fyne.NewSize(400, 400))

	// All the labels and values for input
	labelX := widget.NewLabel("X:")
	valueX := widget.NewEntry()
	labelY := widget.NewLabel("Y:")
	valueY := widget.NewEntry()
	labelZ := widget.NewLabel("Z:")
	valueZ := widget.NewEntry()
	labelMaterial := widget.NewLabel("Material:")
	valueMaterial := widget.NewEntry()
	labelLocation := widget.NewLabel("Location:")
	valueLocation := widget.NewEntry()

	// Create a grid
	grid := container.New(layout.NewGridLayout(2), labelX, valueX, labelY, valueY, labelZ, valueZ,
		labelMaterial, valueMaterial, labelLocation, valueLocation)

	buttonAdd := widget.NewButton("Add Stock", func() {
		x, _ := strconv.Atoi(valueX.Text)
		y, _ := strconv.Atoi(valueY.Text)
		z, _ := strconv.Atoi(valueZ.Text)
		stock := Stock{
			XLength:  x,
			YLength:  y,
			ZLength:  z,
			Material: valueMaterial.Text,
			Location: valueLocation.Text,
		}
		addStock(&stock)
	})

	buttonSearch := widget.NewButton("Search", func() {
		// TODO: Check if x,y,z is an int
		var x, y, z int
		var err error = nil

		x, err = parseIntInput(valueX.Text)
		if err != nil {
			log.Println(err)
			return
		}
		y, err = parseIntInput(valueY.Text)
		if err != nil {
			log.Println(err)
			return
		}
		z, err = parseIntInput(valueZ.Text)
		if err != nil {
			log.Println(err)
			return
		}

		stock := Stock{
			XLength:  x,
			YLength:  y,
			ZLength:  z,
			Material: valueMaterial.Text,
			Location: valueLocation.Text,
		}
		searchStock(&stock)
	})

	// Create a vertical box layout to stack the grid and the button
	content := container.NewVBox(
		grid, buttonSearch, buttonAdd)

	wAdd.SetContent(content)
	wAdd.Show()
}
