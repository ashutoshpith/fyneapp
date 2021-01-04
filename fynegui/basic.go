package fynegui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func Basic(){
	app := app.New()
	window := app.NewWindow("Hello world")
	window.CenterOnScreen()
	window.Resize(fyne.Size{Width: 400, Height: 400})
	


	hello := widget.NewLabel("Hello ")
	window.SetContent(widget.NewVBox(
		hello,
		widget.NewButton("Hi!", func(){
			hello.SetText("Welcome")
		}),
		widget.NewLabel("Hello label"),
	))

	window.ShowAndRun()

}

