package navmenu

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type Item struct {
	Label string
	Path  string
}

func NavMenu(
	items []Item,
	onPress func(string),
) fyne.CanvasObject {
	var menuItems = []fyne.CanvasObject{}

	for _, item := range items {
		menuItems = append(menuItems, createMenuItem(item, onPress))
	}

	bottomBorder := canvas.NewRectangle(color.White)
	bottomBorder.SetMinSize(fyne.NewSize(0, 2))

	itemsList := container.NewHBox(menuItems...)

	content := container.NewBorder(
		itemsList,
		bottomBorder,
		nil,
		nil,
	)

	return content

}

func createMenuItem(item Item, onPress func(string)) fyne.CanvasObject {

	return NewMenuItemWidget(item.Label, func() {
		fmt.Println("Press on", item.Label)
		if onPress != nil {
			onPress(item.Path)
		}
	})

}
