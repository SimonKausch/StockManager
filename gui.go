package main

import (
	"database/sql"
	"fyne.io/fyne/v2/layout"
	"image/color"
	"log"
	"strconv"
	"time"

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

func listStockGUI(db *sql.DB) *string {
	var out string

	slice, err := ListStock(db)
	if err != nil {
		log.Fatal(err)
	}
	for _, s := range slice {
		out += printStock(s)
	}
	return &out
}

func gui() {
	db, err := initDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	a := app.New()
	w := a.NewWindow("StockManger")
	w.Resize(fyne.Size{Width: 750, Height: 500})

	clock := widget.NewLabel("")
	linesep := canvas.NewLine(color.Black)
	allStock := widget.NewLabel("")
	listStockButton := widget.NewButton("List all stock", func() {
		allStock.SetText(*listStockGUI(db))
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
	labelX := widget.NewLabel("X:")
	valueX := widget.NewEntry()
	labelY := widget.NewLabel("Y:")
	valueY := widget.NewEntry()
	labelZ := widget.NewLabel("Y:")
	valueZ := widget.NewEntry()
	labelMaterial := widget.NewLabel("Material:")
	valueMaterial := widget.NewEntry()
	labelLocation := widget.NewLabel("Location:")
	valueLocation := widget.NewEntry()
	grid := container.New(layout.NewGridLayout(2), labelX, valueX, labelY, valueY, labelZ, valueZ,
		labelMaterial, valueMaterial, labelLocation, valueLocation)
	wAdd.SetContent(grid)
	wAdd.Show()
}
