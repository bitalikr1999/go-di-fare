package router

import "fyne.io/fyne/v2"

type component interface {
	Render() fyne.CanvasObject
}

type Resource struct {
	Path      string
	Component component
}

type ResourceList = []Resource
