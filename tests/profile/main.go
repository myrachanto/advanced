package main

import (
	"log"
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	// Create a file to store the profile data
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	defer f.Close()

	// Start CPU profiling
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()

	// Simulate some work for the profile
	time.Sleep(30 * time.Second)
}
