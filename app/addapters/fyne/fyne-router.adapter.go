package adapter_fyne

import (
	uilayouts "bitalikr1999/difare/presentation/layouts"

	"fyne.io/fyne/v2"
)

type FyneRouterAdapter struct {
	window fyne.Window
	layout *uilayouts.MainLayout
}

func NewFyneRouterAdapter(w fyne.Window, layout *uilayouts.MainLayout) *FyneRouterAdapter {
	return &FyneRouterAdapter{
		window: w,
		layout: layout,
	}
}

func (adapter *FyneRouterAdapter) Render(objects fyne.CanvasObject) {
	fyne.Do(func() {
		adapter.window.SetContent(
			adapter.layout.Render([]fyne.CanvasObject{objects}),
		)
	})
}
