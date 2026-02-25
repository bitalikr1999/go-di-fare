package navmenu

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type MenuItemWidget struct {
	widget.BaseWidget
	label *canvas.Text
	onTap func()
}

func NewMenuItemWidget(text string, onTap func()) *MenuItemWidget {
	w := &MenuItemWidget{
		label: canvas.NewText(text, color.White),
		onTap: onTap,
	}
	w.ExtendBaseWidget(w)
	return w
}

func (w *MenuItemWidget) Tapped(*fyne.PointEvent) {
	if w.onTap != nil {
		w.onTap()
	}
}

func (w *MenuItemWidget) CreateRenderer() fyne.WidgetRenderer {

	content := container.NewPadded(
		container.NewCenter(w.label),
	)

	return widget.NewSimpleRenderer(content)
}
