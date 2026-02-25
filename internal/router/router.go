package router

import (
	"errors"
	"fmt"

	"fyne.io/fyne/v2"
)

type Window interface {
	Render(fyne.CanvasObject)
}

type Router struct {
	resources   ResourceList
	defaultPath string
	window      Window
	ctr         *RouterCtr
}

func NewRouter(
	resources ResourceList,
	defaultPath string,
	window Window,
	ctr *RouterCtr,
) *Router {
	var router = &Router{
		resources:   resources,
		defaultPath: defaultPath,
		window:      window,
		ctr:         ctr,
	}

	router.init()

	return router
}

func (r *Router) AddResources(resources ResourceList) {

	for _, resource := range resources {
		r.resources = append(r.resources, resource)
	}

}

func (r *Router) init() {

	if r.defaultPath == "" {
		return
	}

	r.Open(r.defaultPath)
}

func (r *Router) findResource(path string) *Resource {

	var resultItem *Resource

	for _, item := range r.resources {

		if item.Path == path {
			resultItem = &item
			break
		}
	}

	return resultItem

}

func (r *Router) Open(path string) error {

	var resourse = r.findResource(path)
	if resourse == nil {
		return errors.New("Page not found")
	}

	r.window.Render(resourse.Component.Render())

	return nil
}

func (r *Router) Listen() {
	r.ctr.Listen(func(path string) {

		fmt.Print("Pathhhhh listern")
		r.Open(path)
	})
}
