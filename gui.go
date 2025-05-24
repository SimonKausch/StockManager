package main

import (
	"fmt"
	"image/color"
	"log"
	"strconv"

	"fyne.io/fyne/v2/layout"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const stepDir = "../stepReader/files/"

func gui() {
	a := app.New()
	w := a.NewWindow("StockManger")
	w.Resize(fyne.Size{Width: 1000, Height: 750})

	var mainContainer *fyne.Container

	linesep := canvas.NewLine(color.Black)

	entryList := widget.NewLabel("")

	rightSide := analyzeStep()

	buttonListAll := widget.NewButton("List all stock", func() {
		entryList.SetText(*listStockString())
	})
	buttonAdd := widget.NewButton("Add or search Stock", func() {
		addStockWindow(a)
	})

	leftSide := container.NewVBox(buttonListAll, buttonAdd, linesep, entryList)

	mainContainer = container.NewAdaptiveGrid(2, leftSide, rightSide)

	w.SetContent(mainContainer)

	w.ShowAndRun()
}

func analyzeStep() *fyne.Container {
	var cont *fyne.Container

	listFiles, err := getFilesinDir()
	if err != nil {
		output := widget.NewLabel(fmt.Sprint(err))
		cont = container.NewVBox(output)
	} else {
		// var r string
		var file string
		entryResult := widget.NewLabel("")

		output := widget.NewSelect(listFiles, func(r string) {
			file = r
		})
		buttonAnalyze := widget.NewButton("Analyze chosen file", func() {
			if len(file) > 0 {
				bbox := requestBBox(file)

				// From float to a string with a precision of 1 decimal place
				x := fmt.Sprintf("X: %.1f\n", bbox.BoxX)
				y := fmt.Sprintf("X: %.1f\n", bbox.BoxY)
				z := fmt.Sprintf("X: %.1f\n", bbox.BoxZ)

				result := x + y + z
				entryResult.SetText(result)
				// cont = container.NewVBox(outputBox)

			} else {
				entryResult.SetText("Filename empty")
			}
		})
		cont = container.NewVBox(output, buttonAnalyze, entryResult)
	}
	return cont
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
