package main

import (
	"context"
	"encoding/json"
	"os"
	"sync"
	"testing"

	"golang.org/x/sync/semaphore"
)

func createTempFile(t *testing.T, content []Task) string {
	file, err := os.CreateTemp("", "test*.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(content); err != nil {
		t.Fatalf("Failed to write JSON to temp file: %v", err)
	}
	if err := file.Close(); err != nil {
		t.Fatalf("Failed to close temp file: %v", err)
	}
	return file.Name()
}

func TestJsonReader(t *testing.T) {
	// Create temporary JSON files with sample data
	tasks := []Task{
		{Name: "Task1", Age: 30, Role: 1},
		{Name: "Task2", Age: 25, Role: 2},
	}

	filename := createTempFile(t, tasks)
	defer os.Remove(filename) // Clean up

	results := make(chan Task, 2)
	errs := make(chan error, 1)

	var wg sync.WaitGroup
	sem := semaphore.NewWeighted(1)
	ctx := context.Background()

	wg.Add(1)
	if err := sem.Acquire(ctx, 1); err != nil {
		t.Fatalf("Failed to acquire semaphore: %v", err)
	}
	go jsonReader(&wg, sem, filename, results, errs)

	wg.Wait()
	close(results)
	close(errs)

	// Check results
	var receivedTasks []Task
	for task := range results {
		receivedTasks = append(receivedTasks, task)
	}

	if len(receivedTasks) != len(tasks) {
		t.Errorf("Expected %d tasks, got %d", len(tasks), len(receivedTasks))
	}

	// Check errors
	for err := range errs {
		t.Errorf("Unexpected error: %v", err)
	}
}
func TestJsonReader_EmptyFile(t *testing.T) {
	// Create an empty temporary JSON file
	filename := createTempFile(t, []Task{})
	defer os.Remove(filename) // Clean up

	results := make(chan Task, 1)
	errs := make(chan error, 1)

	var wg sync.WaitGroup
	sem := semaphore.NewWeighted(1)
	ctx := context.Background()

	wg.Add(1)
	if err := sem.Acquire(ctx, 1); err != nil {
		t.Fatalf("Failed to acquire semaphore: %v", err)
	}
	go jsonReader(&wg, sem, filename, results, errs)

	wg.Wait()
	close(results)
	close(errs)

	if len(results) > 0 {
		t.Errorf("Unexpected result received")
	}

	if len(errs) == 0 {
		t.Errorf("expected an en empty file error")
	}
}

func TestMainProcess(t *testing.T) {
	// Create temporary JSON files with sample data
	tasks1 := []Task{{Name: "Task1", Age: 30, Role: 1}}
	tasks2 := []Task{{Name: "Task2", Age: 25, Role: 2}}

	file1 := createTempFile(t, tasks1)
	file2 := createTempFile(t, tasks2)
	defer os.Remove(file1) // Clean up
	defer os.Remove(file2) // Clean up

	files := []string{file1, file2}
	workerPoolSize := 2
	results := make(chan Task)
	errs := make(chan error, len(files))

	sem := semaphore.NewWeighted(int64(workerPoolSize))
	ctx := context.Background()

	var wg sync.WaitGroup
	var resWG sync.WaitGroup

	var res []Task
	resWG.Add(1)
	go func() {
		defer resWG.Done()
		for task := range results {
			res = append(res, task)
		}
	}()

	for _, file := range files {
		if err := sem.Acquire(ctx, 1); err != nil {
			t.Fatalf("Failed to acquire semaphore: %v", err)
		}
		wg.Add(1)
		go jsonReader(&wg, sem, file, results, errs)
	}

	wg.Wait()
	close(results)
	resWG.Wait()
	close(errs)

	if len(res) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(res))
	}

	for err := range errs {
		t.Errorf("Unexpected error: %v", err)
	}
}
