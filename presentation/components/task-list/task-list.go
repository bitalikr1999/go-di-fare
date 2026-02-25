package tasklist

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type Item struct {
	Id    int
	Title string
}

func RenderList(items []Item) fyne.CanvasObject {

	tasksComponents := []fyne.CanvasObject{}

	for _, item := range items {
		tasksComponents = append(tasksComponents, RenderTask(item))
	}

	return container.NewVBox(tasksComponents...)

}

func RenderTask(item Item) fyne.CanvasObject {

	text := canvas.NewText(item.Title, color.White)

	bg := canvas.NewRectangle(color.NRGBA{
		R: 40, G: 40, B: 40, A: 255,
	})
	bg.StrokeColor = color.NRGBA{R: 80, G: 80, B: 80, A: 255}
	bg.StrokeWidth = 1

	content := container.NewPadded(text)

	card := container.NewMax(bg, content)

	return card
}
