package main

import (
	"fmt"
	"sync"
	"time"
)

type semaphore struct {
	collecion chan struct{}
}

func New(limit int) *semaphore {
	return &semaphore{
		collecion: make(chan struct{}, limit),
	}
}
func (s *semaphore) Acquire() {
	s.collecion <- struct{}{}
}
func (s *semaphore) Release() {
	<-s.collecion
}
func main() {
	var wg sync.WaitGroup
	sem := New(3)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sem.Acquire()
			defer sem.Release()

			fmt.Printf("Goroutine %d is running\n", i)
			time.Sleep(1 * time.Second)
			fmt.Printf("Goroutine %d is done\n", i)
		}()
	}

	wg.Wait()
}
