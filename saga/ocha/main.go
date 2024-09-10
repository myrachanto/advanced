package main

import (
	"errors"
	"fmt"
)

// Saga Step Interface
type SagaStep interface {
	Execute() error
	Compensate() error
}

// Payment Step
type PaymentStep struct{}

func (p PaymentStep) Execute() error {
	fmt.Println("Executing Payment Step")
	// Simulate success
	return nil
}

func (p PaymentStep) Compensate() error {
	fmt.Println("Compensating Payment Step")
	// Simulate compensation logic
	return nil
}

// Inventory Step
type InventoryStep struct{}

func (i InventoryStep) Execute() error {
	fmt.Println("Executing Inventory Step")
	// Simulate failure
	return errors.New("inventory failed")
}

func (i InventoryStep) Compensate() error {
	fmt.Println("Compensating Inventory Step")
	// Simulate compensation logic
	return nil
}

// Orchestrator
type SagaOrchestrator struct {
	steps []SagaStep
}

func NewSagaOrchestrator(steps ...SagaStep) *SagaOrchestrator {
	return &SagaOrchestrator{steps: steps}
}

func (s *SagaOrchestrator) Execute() {
	for i, step := range s.steps {
		err := step.Execute()
		if err != nil {
			fmt.Println("Error:", err)
			// Rollback if any step fails
			for j := i - 1; j >= 0; j-- {
				_ = s.steps[j].Compensate()
			}
			return
		}
	}
	fmt.Println("Saga Completed Successfully")
}

func main() {
	saga := NewSagaOrchestrator(
		PaymentStep{},
		InventoryStep{},
	)
	saga.Execute()
}
