package view

import (
	"groupietracker/model"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var (
	myApp model.AppData
)

func Start2() {
	myApp.App = app.New()
	myApp.Window = myApp.App.NewWindow("Groupie Tracker.")
	PageMain()
	myApp.Window.Resize(fyne.NewSize(1, 750))
	myApp.Window.ShowAndRun()
}
