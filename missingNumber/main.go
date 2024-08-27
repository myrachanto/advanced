package main

import "fmt"

func missingNumber(numbers []int) int {
	l := len(numbers) + 1
	expectedSum := l * (l + 1) / 2
	actual := 0
	for _, i := range numbers {
		actual += i
	}
	return expectedSum - actual
}
func missingNumber2(numbers []int) int {
	// Calculate first term and difference of the sequence
	firstTerm := numbers[0]
	n := len(numbers) + 1
	lastTerm := firstTerm + (n-1)*2 // last term in the complete sequence

	// Calculate expected sum of the arithmetic sequence
	expectedSum := n * (firstTerm + lastTerm) / 2

	// Calculate actual sum of given numbers
	actualSum := 0
	for _, num := range numbers {
		actualSum += num
	}

	return expectedSum - actualSum
}
func main() {
	list := []int{5, 10, 15, 25}
	// 5*6 = 30/2 =15
	fmt.Println(missingNumber2(list))

}
