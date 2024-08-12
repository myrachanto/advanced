package main

import (
	"container/list"
	"fmt"
)

type MyCache struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}
type entry struct {
	key   int
	value int
}

func New(cap int) *MyCache {
	if cap == 0 {
		cap = 50
	}
	return &MyCache{
		capacity: cap,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}
func (c *MyCache) Get(key int) int {
	if e, ok := c.cache[key]; ok {
		c.list.MoveToFront(e)
		return e.Value.(*entry).value
	}
	return -1
}
func (c *MyCache) Put(data entry) {
	if e, ok := c.cache[data.key]; ok {
		c.list.MoveToFront(e)
		e.Value.(*entry).value = data.value
	} else {
		if c.list.Len() >= c.capacity {
			tail := c.list.Back()
			if tail != nil {
				delete(c.cache, tail.Value.(*entry).key)
				c.list.Remove(tail)
			}
		}
		el := c.list.PushFront(&data)
		c.cache[data.key] = el
	}
}
func main() {
	cache := New(2)
	cache.Put(entry{1, 1})
	cache.Put(entry{2, 2})
	fmt.Println(cache.Get(1)) // returns 1
	cache.Put(entry{3, 3})    // evicts key 2
	fmt.Println(cache.Get(2)) // returns -1 (not found)
	cache.Put(entry{4, 4})    // evicts key 1
	fmt.Println(cache.Get(1)) // returns -1 (not found)
	fmt.Println(cache.Get(3)) // returns 3
	fmt.Println(cache.Get(4)) // returns 4
}
