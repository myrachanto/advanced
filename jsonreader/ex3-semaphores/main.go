package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
)

type Task struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
	Role int    `json:"role,omitempty"`
}

func jsonReader(filename string, results chan Task) error {

	file, err := os.Open(filename) // Change the filename as needed
	if err != nil {
		fmt.Println("could not open the file")
		return err
	}
	// Ensure the file is closed after processing
	// Move this line right after opening the file
	defer file.Close()
	// Check if the file is empty
	stat, err := file.Stat()
	if err != nil {
		return err
	}
	if stat.Size() == 0 {
		return fmt.Errorf("file is empty %s", filename)
	}

	decoder := json.NewDecoder(file)
	var tasks []Task
	if err := decoder.Decode(&tasks); err != nil {
		if err == io.EOF {
			// Ignore EOF errors for empty files
			return nil
		}
		return err
	}
	for _, task := range tasks {
		results <- task
	}
	return nil
}

func main() {
	files := []string{"./jsonreader/ex3-semaphores/data/file.json", "./jsonreader/ex2/data/file2.json", "./jsonreader/ex2/data/file3.json", "/jsonreader/ex2/data/file4.json"}
	workerpool := 2
	results := make(chan Task, 50)
	sem := semaphore.NewWeighted(int64(workerpool))
	errg, ctx := errgroup.WithContext(context.Background())
	for i := range len(files) {
		if err := sem.Acquire(ctx, 1); err != nil {
			log.Fatal(err)
		}
		errg.Go(func() error {
			defer sem.Release(1)
			return jsonReader(files[i], results)
		})

	}
	go func() {
		err := errg.Wait()
		if err != nil {
			log.Fatal(err)
		}
		close(results)
	}()
	var res []Task
	for task := range results {
		res = append(res, task)
	}
	fmt.Println(len(res))

}
