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

func main() {
	list := []int{1, 2, 3, 5}
	// 5*6 = 30/2 =15
	fmt.Println(missingNumber(list))

}
