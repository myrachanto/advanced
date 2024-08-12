package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	"golang.org/x/sync/semaphore"
)

type Task struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
	Role int    `json:"role,omitempty"`
}

func jsonReader(wg *sync.WaitGroup, sem *semaphore.Weighted, filename string, results chan Task, errs chan error) {
	fmt.Printf("working on %s \n", filename)
	defer wg.Done()
	defer sem.Release(1)
	file, err := os.Open(filename) // Change the filename as needed
	if err != nil {
		errs <- err
		return
	}
	// Ensure the file is closed after processing
	// Move this line right after opening the file
	defer file.Close()
	// Check if the file is empty
	stat, err := file.Stat()
	if err != nil {
		errs <- err
		return
	}
	if stat.Size() == 0 {
		errs <- fmt.Errorf("file %s is empty", filename)
		return
	}

	decoder := json.NewDecoder(file)
	var tasks []Task
	if err := decoder.Decode(&tasks); err != nil {
		if err == io.EOF {
			errs <- fmt.Errorf("failed to decode JSON from file %s: %w", filename, err)
			return
		}
		return

	}
	for _, task := range tasks {
		results <- task
	}
}

func main() {
	files := []string{"./jsonreader/ex3-semaphores/data/file.json", "./jsonreader/ex3-semaphores/data/file2.json", "./jsonreader/ex3-semaphores/data/file3.json", "/jsonreader/ex3-semaphores/data/file4.json"}
	workerpool := 2
	results := make(chan Task)
	// results := make(chan Task, 50) // when you know your record size
	errs := make(chan error, len(files))

	sem := semaphore.NewWeighted(int64(workerpool))
	ctx := context.Background()

	var wg sync.WaitGroup
	var wg1 sync.WaitGroup

	// Goroutine to collect results without knowing the buffer size

	var res []Task

	wg1.Add(1)
	go func() {
		defer wg1.Done()
		for task := range results {
			res = append(res, task)
		}
		fmt.Printf("Number of tasks processed: %d\n", len(res))
	}()

	for _, file := range files {
		if err := sem.Acquire(ctx, 1); err != nil {
			log.Fatal(err)
		}
		wg.Add(1)
		go jsonReader(&wg, sem, file, results, errs)

	}
	wg.Wait()
	close(results)
	wg1.Wait()
	close(errs)

	fmt.Println(len(res))
	for err := range errs {
		fmt.Println(err)
	}

}
