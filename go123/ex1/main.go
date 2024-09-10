package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	// reversing the order slice
	s := []int{1, 2, 3, 4, 5}
	for i, x := range slices.Backward(s) {
		fmt.Println(i, x)
	}
	//
	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
	m2 := maps.Collect(maps.All(m1))
	fmt.Println(m2) // Output: map[a:1 b:2 c:3]

	// Deleting a item from a slice
	nums := []int{1, 2, 3, 4, 5, 6}
	evenNumbersRemoved := slices.DeleteFunc(nums, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println(evenNumbersRemoved) // Output: [1, 3, 5]
	//
	words := []string{"go", "go", "is", "great", "great"}
	uniqueWords := slices.CompactFunc(words, func(a, b string) bool {
		return a == b
	})
	fmt.Println(uniqueWords) // Output: [go is great]
	fmt.Println(slices.IsSorted([]string{"Alice", "Bob", "Vera"}))
	numbers := []int{0, 42, -10, 8}
	fmt.Println(slices.Min(numbers))

	m1s := map[string]int{"a": 1, "b": 2}
	m2s := map[string]int{"a": 1, "b": 2}
	areEqual := maps.EqualFunc(m1s, m2s, func(a, b int) bool {
		return a == b
	})
	fmt.Println(areEqual) // Output: true

	names := []string{"Alice", "Bob", "Vera"}
	for i, v := range slices.All(names) {
		fmt.Println(i, ":", v)
	}
	seq := func(yield func(int) bool) {
		for i := 0; i < 10; i += 2 {
			if !yield(i) {
				return
			}
		}
	}

	ss := slices.AppendSeq([]int{1, 2}, seq)
	fmt.Println(ss)

	names2 := []string{"Alice", "Bob", "Vera"}
	n, found := slices.BinarySearch(names2, "Vera")
	fmt.Println("Vera:", n, found)
	n, found = slices.BinarySearch(names2, "Bill")
	fmt.Println("Bill:", n, found)

	ssd := []int{1, 2, 3, 4, 5, 6, 7}
	for c := range slices.Chunk(ssd, 2) {
		fmt.Println(c)
	}

	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sl := a[:4]
	clip := slices.Clip(sl)
	fmt.Println(cap(sl))
	fmt.Println(clip)
	fmt.Println(len(clip))
	fmt.Println(cap(clip))

	letters := []string{"a", "b", "c", "d", "e"}
	letters = slices.Delete(letters, 1, 4)
	fmt.Println(letters)

	names3 := []string{"Alice", "Bob", "Vera"}
	names3 = slices.Insert(names3, 1, "Bill", "Billie")
	names3 = slices.Insert(names3, len(names3), "Zac")
	fmt.Println(names3)

}
