package main

import (
	"fmt"
	"image/color"
	"log"
	"strconv"

	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const stepDir = "../stepReader/files/"

func gui() {
	a := app.NewWithID("stepReader")
	w := a.NewWindow("StockManger")
	w.Resize(fyne.Size{Width: 750, Height: 750})

	Linesep := canvas.NewLine(color.Black)

	entryList := widget.NewLabel("")

	rightSide := analyzeStep(w)

	buttonListAll := widget.NewButton("List all stock", func() {
		temporaryStock, _ := ListStock()
		entryList.SetText(createTable(temporaryStock))

		// FIX: Delete later, used for debugging
		l, _ := ListByMaterial("Aluminium")
		log.Println(l)
	})

	buttonListByMaterial := widget.NewButton("List stock by material", func() {
		listByMaterialWindow(a)
	})

	buttonAdd := widget.NewButton("Add or search Stock", func() {
		addStockWindow(a)
	})
	buttonRemove := widget.NewButton("Remove Stock", func() {
		removeStockWindow(a)
	})

	leftSide := container.NewVBox(buttonListAll, buttonListByMaterial, buttonAdd, buttonRemove, Linesep, entryList)

	mainContainer := container.NewAdaptiveGrid(2, leftSide, rightSide)

	w.SetContent(mainContainer)
	w.CenterOnScreen()
	w.ShowAndRun()
}

func analyzeStep(w fyne.Window) *fyne.Container {
	var cont *fyne.Container
	var bbox BoundingBox
	var err error

	// Create a label to display the selected file path
	selectedFilePathLabel := widget.NewLabel("No file selected")

	// Open file dialog
	buttonFile := widget.NewButton("Choose .stp file", func() {
		fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, w)
			}
			if reader == nil {
				selectedFilePathLabel.SetText("File selection cancelled")
			}
			// Get the full file path
			selectedFilePathLabel.SetText(reader.URI().Path())

			defer reader.Close()
		}, w)

		// Filte for stp files
		fd.SetFilter(storage.NewExtensionFileFilter([]string{".stp"}))

		fd.Show()
	})

	// Label to display results
	entryResult := widget.NewLabel("")

	// Analyze step file selected with file dialog
	buttonAnalyze := widget.NewButton("Get bounding box of step file", func() {
		if len(selectedFilePathLabel.Text) > 0 {
			bbox, err = uploadFile(selectedFilePathLabel.Text)
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
			dialog.ShowInformation("Info", "Choose file first", w)
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

	cont = container.NewVBox(selectedFilePathLabel, buttonFile, buttonAnalyze, buttonFindStock, canvas.NewLine(color.Black), entryResult)

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
	selectMaterial := widget.NewSelect(listMaterials, func(selectedMat string) {
	})
	valueMaterial := widget.NewEntry()
	labelLocation := widget.NewLabel("Location:")
	valueLocation := widget.NewEntry()

	// Create a grid
	grid := container.New(layout.NewGridLayout(2), labelX, valueX, labelY, valueY, labelZ, valueZ,
		labelMaterial, selectMaterial, labelLocation, valueLocation)

	buttonAdd := widget.NewButton("Add Stock", func() {
		x, _ := strconv.Atoi(valueX.Text)
		y, _ := strconv.Atoi(valueY.Text)
		z, _ := strconv.Atoi(valueZ.Text)
		stock := Stock{
			XLength:  x,
			YLength:  y,
			ZLength:  z,
			Material: selectMaterial.Selected,
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

func listByMaterialWindow(a fyne.App) {
	wAdd := a.NewWindow("Filter by material")
	wAdd.Resize(fyne.NewSize(400, 400))

	// All the labels and values for input
	allMaterials, err := ListMaterials()
	if err != nil {
		log.Println(err)
	}
	labelMaterial := widget.NewLabel("Material:")
	selectMaterial := widget.NewSelect(allMaterials, func(selectedMat string) {
	})

	// Create a grid
	grid := container.New(layout.NewGridLayout(2), labelMaterial, selectMaterial)

	// Show output from search
	resultsList := widget.NewLabel("")

	// TODO: Search by chosen material
	buttonByMaterial := widget.NewButton("Search by material", func() {
		res, err := ListByMaterial(selectMaterial.Selected)
		if err != nil {
			log.Println(err)
		}
		resultsList.SetText(createTable(res))
	})

	// Create a vertical box layout to stack the grid and the button
	content := container.NewVBox(
		grid, buttonByMaterial, resultsList)

	wAdd.SetContent(content)
	wAdd.Show()
}
