package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	list := map[string]int{
		"a": 1,
		"d": 4,
		"b": 2,
		"e": 5,
		"c": 3,
	}
	fmt.Println(slices.Sorted(maps.Keys(list)))
	fmt.Println(slices.Sorted(maps.Values(list)))
	res := []int{}
	for _, g := range slices.Backward(slices.Sorted(maps.Values(list))) {
		res = append(res, g)
	}
	fmt.Println(res)
}
