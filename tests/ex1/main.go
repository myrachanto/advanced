package main

import (
	"fmt"
	"slices"
)

func min(data ...int64) int64 {
	if len(data) == 0 {
		// Handle the case where no numbers are provided
		return 0 // or an error value if you prefer
	}
	leastnumber := int64(1<<63 - 1)
	for _, v := range data {
		if v < leastnumber {
			leastnumber = v
		}
	}
	return leastnumber
}

func main() {
	list := []int64{61, 53, 71, 76, 97, 34, 32, 45, 16, 45}
	fmt.Println(min(list...))
	fmt.Println(slices.Min(list))

}
