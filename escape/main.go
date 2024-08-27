package main

import "fmt"

func newInt() *int {
	i := 42
	return &i // `i` escapes to heap because it's returned as a pointer
}
func closure() func() {
	s := "hello"
	return func() { // `s` escapes to heap because it's captured in a closure
		fmt.Println(s)
	}
}
func startGoroutine() {
	s := "goroutine"
	go func() { // `s` escapes to heap because it's used in a goroutine
		fmt.Println(s)
	}()
}
func main() {
	newInt()
	closure()
	startGoroutine()
}
