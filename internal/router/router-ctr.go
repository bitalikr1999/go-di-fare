package router

import "fmt"

type RouterCtr struct {
	navigateCh chan string
}

func NewRouterCtr() *RouterCtr {
	return &RouterCtr{
		navigateCh: make(chan string),
	}
}

func (r *RouterCtr) Listen(callback func(string)) {
	for path := range r.navigateCh {
		fmt.Print("On TEST")
		callback(path)
	}
}

func (r *RouterCtr) Close() {
	close(r.navigateCh)
}

func (r *RouterCtr) Navigate(path string) {
	r.navigateCh <- path
}
