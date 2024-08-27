package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string)
	for _, str := range strs {
		sortedStr := sortString(str)
		anagrams[sortedStr] = append(anagrams[sortedStr], str)
	}
	results := [][]string{}
	for _, res := range anagrams {
		results = append(results, res)
	}
	return results
}

func sortString(s string) string {
	slice := []rune(s)
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] > slice[j]
	})
	return string(slice)
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	sort.Strings(strs)
	fmt.Println("sorted strings:", strs)

	fmt.Println(groupAnagrams(strs))
}
