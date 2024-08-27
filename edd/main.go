package main

import (
	"fmt"
	"time"
)

// Event represents a generic event with a type and payload.
type Event struct {
	Type      string
	Timestamp time.Time
	Payload   interface{}
}

func main() {
	bus := New()

	// Subscribe to "user:created" event
	bus.Subscribe("user:created", func(event Event) {
		fmt.Printf("User created event received at %s: %v\n", event.Timestamp, event.Payload)
	})

	// Subscribe to "order:placed" event
	bus.Subscribe("order:placed", func(event Event) {
		fmt.Printf("Order placed event received at %s: %v\n", event.Timestamp, event.Payload)
	})

	// Publish a "user:created" event
	bus.Publish(Event{
		Type:      "user:created",
		Timestamp: time.Now(),
		Payload:   "User123",
	})

	// Publish an "order:placed" event
	bus.Publish(Event{
		Type:      "order:placed",
		Timestamp: time.Now(),
		Payload:   "Order456",
	})

	// Wait to ensure all events are processed
	time.Sleep(1 * time.Second)
}
