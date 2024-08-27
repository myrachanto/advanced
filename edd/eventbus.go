package main

import "sync"

type EventHandler func(event Event)

type EventBus struct {
	subscriber map[string][]EventHandler
	mu         sync.RWMutex
}

func New() *EventBus {
	return &EventBus{
		subscriber: make(map[string][]EventHandler),
	}
}

func (bus *EventBus) Subscribe(eventType string, handler EventHandler) {
	bus.mu.Lock()
	defer bus.mu.Unlock()
	bus.subscriber[eventType] = append(bus.subscriber[eventType], handler)
}
func (bus *EventBus) Publish(event Event) {
	bus.mu.RLock()
	defer bus.mu.RUnlock()
	if handlers, ok := bus.subscriber[event.Type]; ok {
		for _, handler := range handlers {
			go handler(event)
		}
	}
}
