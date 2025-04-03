package main

import (
	"image/color"
	"log"
	"strconv"
	"time"

	"fyne.io/fyne/v2/layout"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

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

	clock := widget.NewLabel("")
	linesep := canvas.NewLine(color.Black)
	allStock := widget.NewLabel("")
	listStockButton := widget.NewButton("List all stock", func() {
		allStock.SetText(*listStockGUI())
	})
	addStockButton := widget.NewButton("Add Stock", func() {
		addStockWindow(a)
	})

	updateTime(clock)
	w.SetContent(clock)

	w.SetContent(container.NewVBox(clock, linesep, allStock, listStockButton, addStockButton))

	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()

	w.ShowAndRun()
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
	valueMaterial := widget.NewEntry()
	labelLocation := widget.NewLabel("Location:")
	valueLocation := widget.NewEntry()

	// Create a grid
	grid := container.New(layout.NewGridLayout(2), labelX, valueX, labelY, valueY, labelZ, valueZ,
		labelMaterial, valueMaterial, labelLocation, valueLocation)

	button := widget.NewButton("Add Stock", func() {
		// Handle button action
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
		addStock(stock)
	})

	// Create a vertical box layout to stack the grid and the button
	content := container.NewVBox(
		grid,
		container.NewCenter(
			button,
		),
	)

	// button := container.New(layout.NewCenterLayout(), widget.NewLabel("TEST"))
	wAdd.SetContent(content)
	// wAdd.SetContent(button)
	wAdd.Show()
}
