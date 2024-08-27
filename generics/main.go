package main

import (
	"fmt"
)

type gen interface {
	int | int64 | uint64 | float32 | float64
}

func Max[T gen](data []T) T {
	if len(data) == 0 {
		return 0
	}
	max := data[0]
	for _, g := range data {
		if max < g {
			max = g
		}

	}
	return max
}
func main() {
	intTests := []int{1, 2, 3, 4, 5, 6, 7}
	floatTests := []float64{1.2, 2.2, 0.2, 0, 4.5}
	mixedTests := []int{-2, -3, 4, 5, 1, 0}

	fmt.Printf("The max number is %v\n", Max(intTests))
	fmt.Printf("The max number is %v\n", Max(floatTests))
	fmt.Printf("The max number is %v\n", Max(mixedTests))

	// interesting concept
	// smallSlice := append([]int(nil), largeSlice[start:end]...)

}
