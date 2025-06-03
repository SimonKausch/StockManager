package main

import (
	"fmt"
	"image/color"
	"log"
	"strconv"

	"fyne.io/fyne/v2/dialog"
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
	w.Resize(fyne.Size{Width: 750, Height: 750})

	Linesep := canvas.NewLine(color.Black)

	entryList := widget.NewLabel("")

	rightSide := analyzeStep(w)

	buttonListAll := widget.NewButton("List all stock", func() {
		temporaryStock, _ := ListStock()
		entryList.SetText(createTable(temporaryStock))
	})
	buttonAdd := widget.NewButton("Add or search Stock", func() {
		addStockWindow(a)
	})
	buttonRemove := widget.NewButton("Remove Stock", func() {
		removeStockWindow(a)
	})

	leftSide := container.NewVBox(buttonListAll, buttonAdd, buttonRemove, Linesep, entryList)

	mainContainer := container.NewAdaptiveGrid(2, leftSide, rightSide)

	w.SetContent(mainContainer)
	w.CenterOnScreen()
	w.ShowAndRun()
}

func analyzeStep(w fyne.Window) *fyne.Container {
	var cont *fyne.Container
	var bbox BoundingBox

	listFiles, err := getFilesinDir()

	if err != nil {
		dialog.ShowError(err, w)
	} else {
		var file string
		entryResult := widget.NewLabel("")

		output := widget.NewSelect(listFiles, func(r string) {
			file = r
		})
		output.PlaceHolder = "Select step file"

		buttonAnalyze := widget.NewButton("Get bounding box of step file", func() {
			if len(file) > 0 {
				bbox, err = requestBBox(file)
				if err != nil {
					dialog.ShowError(err, w)
				}

				// From float to a string with a precision of 1 decimal place
				x := fmt.Sprintf("X: %.1f\n", bbox.BoxX)
				y := fmt.Sprintf("X: %.1f\n", bbox.BoxY)
				z := fmt.Sprintf("X: %.1f\n", bbox.BoxZ)

				result := x + y + z
				entryResult.SetText(result)
			} else {
				dialog.ShowInformation("Info", "Filename empty", w)
			}
		})

		buttonFindStock := widget.NewButton("Find fitting stock", func() {
			if entryResult.Text == "" {
				dialog.ShowInformation("Error", "No step file analyzed", w)
			} else {
				allStock, err := ListStock()
				if err != nil {
					log.Println(err)
				}
				fittingStock := findFittingStock(bbox, allStock)
				log.Println(fittingStock)

				// TODO: Instead of dialog, create new window
				dialog.ShowInformation("Fitting stock", createTable(fittingStock), w)
			}
		})

		cont = container.NewVBox(output, buttonAnalyze, buttonFindStock, canvas.NewLine(color.Black), entryResult)
	}
	return cont
}

func removeStockWindow(a fyne.App) {
	wAdd := a.NewWindow("Remove stock")
	wAdd.Resize(fyne.NewSize(300, 200))

	labelID := widget.NewLabel("ID:")
	valueID := widget.NewEntry()
	buttonRemove := widget.NewButton("Remove stock", func() {
		id, err := parseIntInput(valueID.Text)
		if err != nil {
			log.Println(err)
		}
		removeStock(id)
	})
	output := widget.NewLabel("")

	grid := container.New(layout.NewGridLayout(2), labelID, valueID)
	content := container.NewVBox(grid, buttonRemove, output)

	wAdd.SetContent(content)
	wAdd.Show()
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
	// TODO: Dropdown with all materials
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

	// Show output from search
	resultsList := widget.NewLabel("")

	buttonSearch := widget.NewButton("Search", func() {
		var x, y, z int
		var err error = nil

		x, err = parseIntInput(valueX.Text)
		if err != nil {
			resultsList.SetText(fmt.Sprint(err))
			return
		}
		y, err = parseIntInput(valueY.Text)
		if err != nil {
			resultsList.SetText(fmt.Sprint(err))
			return
		}
		z, err = parseIntInput(valueZ.Text)
		if err != nil {
			resultsList.SetText(fmt.Sprint(err))
			return
		}

		stock := Stock{
			XLength:  x,
			YLength:  y,
			ZLength:  z,
			Material: valueMaterial.Text,
			Location: valueLocation.Text,
		}

		// Display result of search or error
		result, err := searchStock(&stock)
		if err != nil {
			resultsList.SetText(fmt.Sprint(err))
		}
		resultsList.SetText("Found suitable stock:\n" + createTable(result))
	})

	// Create a vertical box layout to stack the grid and the button
	content := container.NewVBox(
		grid, buttonSearch, buttonAdd, resultsList)

	wAdd.SetContent(content)
	wAdd.Show()
}
