package uilayouts

import (
	navmenu "bitalikr1999/difare/presentation/components/nav-menu"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type MainLayout struct {
	Items       []navmenu.Item
	OnClickItem func(string)
}

func (l MainLayout) Render(
	components []fyne.CanvasObject,
) fyne.CanvasObject {

	root := container.NewBorder(
		navmenu.NavMenu(l.Items, l.OnClickItem),
		nil,
		nil,
		nil,
		components...,
	)
	return root

}
