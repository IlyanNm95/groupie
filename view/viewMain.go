package view

import (
	"groupietracker/model"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var (
	myApp model.AppData
)

func Start() {
	myApp.App = app.New()
	myApp.Window = myApp.App.NewWindow("Groupie Tracker.")
	PageMain()
	myApp.Window.Resize(fyne.NewSize(225, 875))
}
