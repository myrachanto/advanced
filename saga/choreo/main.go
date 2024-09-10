package main

import (
	"fmt"
	"sync"
	"time"
)

// Event types
type EventType string

const (
	OrderCreated      EventType = "OrderCreated"
	PaymentSuccess    EventType = "PaymentSuccess"
	PaymentFailed     EventType = "PaymentFailed"
	InventoryDeducted EventType = "InventoryDeducted"
	InventoryFailed   EventType = "InventoryFailed"
	OrderCanceled     EventType = "OrderCanceled"
)

// Event structure
type Event struct {
	Type EventType
	Data string
}

// Channels for events
var paymentChannel = make(chan Event)
var inventoryChannel = make(chan Event)
var orderChannel = make(chan Event)

func orderService(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Order Service: Creating Order...")
	time.Sleep(1 * time.Second)
	// Order created event
	orderChannel <- Event{Type: OrderCreated, Data: "Order123"}
}

func paymentService(wg *sync.WaitGroup) {
	defer wg.Done()

	for event := range orderChannel {
		if event.Type == OrderCreated {
			fmt.Println("Payment Service: Processing Payment for", event.Data)
			time.Sleep(1 * time.Second)

			// Simulate payment success or failure
			success := true
			if success {
				fmt.Println("Payment Service: Payment Successful")
				paymentChannel <- Event{Type: PaymentSuccess, Data: event.Data}
			} else {
				fmt.Println("Payment Service: Payment Failed")
				paymentChannel <- Event{Type: PaymentFailed, Data: event.Data}
			}
		}
	}
}

func inventoryService(wg *sync.WaitGroup) {
	defer wg.Done()

	for event := range paymentChannel {
		if event.Type == PaymentSuccess {
			fmt.Println("Inventory Service: Deducting Inventory for", event.Data)
			time.Sleep(1 * time.Second)

			// Simulate inventory success or failure
			success := true
			if success {
				fmt.Println("Inventory Service: Inventory Deducted")
				inventoryChannel <- Event{Type: InventoryDeducted, Data: event.Data}
			} else {
				fmt.Println("Inventory Service: Inventory Deduction Failed")
				inventoryChannel <- Event{Type: InventoryFailed, Data: event.Data}
			}
		} else if event.Type == PaymentFailed {
			fmt.Println("Inventory Service: Payment Failed, Canceling Order", event.Data)
			orderChannel <- Event{Type: OrderCanceled, Data: event.Data}
		}
	}
}

func orderSagaListener(wg *sync.WaitGroup) {
	defer wg.Done()

	for event := range inventoryChannel {
		if event.Type == InventoryDeducted {
			fmt.Println("Order Saga: Order Completed Successfully for", event.Data)
		} else if event.Type == InventoryFailed {
			fmt.Println("Order Saga: Inventory Failed, Canceling Order", event.Data)
			orderChannel <- Event{Type: OrderCanceled, Data: event.Data}
		}
	}

	// Handle order cancelation events
	for event := range orderChannel {
		if event.Type == OrderCanceled {
			fmt.Println("Order Saga: Order Canceled for", event.Data)
		}
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(4)
	go orderService(&wg)
	go paymentService(&wg)
	go inventoryService(&wg)
	go orderSagaListener(&wg)

	wg.Wait()
	fmt.Println("Saga Completed.")
}
