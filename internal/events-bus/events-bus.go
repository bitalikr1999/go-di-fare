package eventsbus

import (
	"fmt"
	"sync"
)

type EventsBus struct {
	subscribers map[string][]chan interface{}
	mu          sync.RWMutex
}

func New() *EventsBus {

	return &EventsBus{
		subscribers: make(map[string][]chan interface{}),
	}
}

func (bus *EventsBus) Subscribe(event string) <-chan interface{} {
	ch := make(chan interface{})

	bus.mu.Lock()
	bus.subscribers[event] = append(bus.subscribers[event], ch)
	bus.mu.Unlock()

	return ch
}

func (bus *EventsBus) Publish(event string, payload interface{}) {

	bus.mu.RLock()
	subs, ok := bus.subscribers[event]
	bus.mu.RUnlock()

	if !ok {
		return
	}

	for _, subChain := range subs {

		go func(ch chan interface{}) {

			select {
			case ch <- payload:
			default:
				fmt.Println("Send event failed")
			}

		}(subChain)
	}
}

func (bus *EventsBus) Unsubscribe(event string) {

	bus.mu.Lock()
	defer bus.mu.Unlock()

	subs, ok := bus.subscribers[event]
	if !ok {
		return
	}

	for _, sub := range subs {
		close(sub)
	}

	delete(bus.subscribers, event)

}

func (bus *EventsBus) Close() {

	bus.mu.Lock()
	defer bus.mu.Unlock()

	for key, eventGroup := range bus.subscribers {
		for _, chain := range eventGroup {
			close(chain)
		}
		delete(bus.subscribers, key)
	}

}
