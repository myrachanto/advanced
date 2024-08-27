package main

import (
	"fmt"
)

func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[[26]int][]string)

	for _, word := range strs {
		// Create a frequency array for the word
		var freq [26]int
		for _, char := range word {
			fmt.Println("========", char-'a', 'a')
			freq[char-'a']++
		}
		// Use frequency array as key
		fmt.Println("========>>>>>", anagrams[freq])
		anagrams[freq] = append(anagrams[freq], word)
	}

	// Collect grouped anagrams
	result := [][]string{}
	for _, group := range anagrams {
		result = append(result, group)
	}
	return result
}

func main() {
	words := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result := groupAnagrams(words)
	fmt.Println(result)
}
