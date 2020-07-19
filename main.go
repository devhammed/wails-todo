package main

import (
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"

	"wails-todo/models"
)

func main() {
	js := mewn.String("./frontend/build/static/js/main.js")
	css := mewn.String("./frontend/build/static/css/main.css")
	app := wails.CreateApp(&wails.AppConfig{
		Width:     600,
		Height:    600,
		Title:     "Wails Todo",
		JS:        js,
		CSS:       css,
		Resizable: true,
		Colour:    "#f1f1f1",
	})

	app.Bind(models.NewTodos())
	app.Run()
}
