package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Ensure Done is called to avoid deadlocks

	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	// Incorrect placement of wg.Add()
	go worker(1, &wg)
	wg.Add(1) // This should be before starting the goroutine

	// Missed Done() call: Uncomment the following lines to simulate
	/*
		wg.Add(1)
		go func() {
			fmt.Println("Anonymous worker starting")
			// Forgetting to call wg.Done() here will cause the program to hang
		}()
	*/

	// Overuse of Add()
	wg.Add(2) // Adds more than required

	go worker(2, &wg)

	wg.Wait() // This will wait indefinitely due to the overuse of Add()
	fmt.Println("All workers completed")
}
