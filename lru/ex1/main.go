package main

import (
	"container/list"
	"fmt"
)

func main() {

	listing := list.New()
	listing.PushBack("one")
	listing.PushBack("two")
	listing.PushBack("three")
	listing.PushBack("four")

	// Adding elements to the front
	listing.PushFront("zero")

	// Adding elements before and after a specific element
	firstElement := listing.PushBack("three")
	listing.InsertBefore("twoandhalf", firstElement)
	listing.InsertAfter("threepoitfive", firstElement)
	// printing elements in a list
	for element := listing.Front(); element != nil; element = element.Next() {
		fmt.Println(element)
	}
}
