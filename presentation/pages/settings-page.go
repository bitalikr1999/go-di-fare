package pages

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

type SettingsPage struct{}

func (page *SettingsPage) Render() fyne.CanvasObject {

	title := canvas.NewText("Settings", color.White)

	content := container.New(layout.NewHBoxLayout(), title)

	return content
}
