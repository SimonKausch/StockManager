package main

import (
	"image/color"
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

// List all stock in db
// allStock, err := ListStock(db)
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	for _, stock := range allStock {
//		fmt.Println(stock)
//	}
func listStockGUI() *string {
	s := "Result of query"
	return &s
}

func gui() {
	var width float32 = 500

	a := app.New()
	w := a.NewWindow("StockManger")
	w.Resize(fyne.Size{Width: width, Height: 200})

	clock := widget.NewLabel("")
	linesep := canvas.NewLine(color.Black)
	allStock := widget.NewLabel("")
	listStockButton := widget.NewButton("List all stock", func() {
		allStock.SetText(*listStockGUI())
	})

	updateTime(clock)
	w.SetContent(clock)

	w.SetContent(container.NewVBox(clock, linesep, allStock, listStockButton))

	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()

	w.ShowAndRun()
}
