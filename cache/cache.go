package cache

import "fmt"

type LRU struct {
	Queue     []int
	Capacity  int
	cache_map map[int]int
}

func indexOf(element int, data []int) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}

func New() *LRU {
	return &LRU{
		Queue:     make([]int, 0),
		Capacity:  5,
		cache_map: make(map[int]int),
	}
}

func (c *LRU) Add(el int) {
	if _, ok := c.cache_map[el]; ok {
		index := indexOf(el, c.Queue)
		if index != -1 {
			c.Queue = append(c.Queue[:index], c.Queue[index+1:]...)
			c.Queue = append([]int{el}, c.Queue...)
			c.cache_map[el] = el
		}
	} else {
		c.cache_map[el] = el
		c.Queue = append([]int{el}, c.Queue...)
		if len(c.Queue) > 5 {
			c.Queue = c.Queue[0:5]
		}
	}
}

func (c *LRU) PrintQueue() {
	fmt.Println(c.Queue)
}
