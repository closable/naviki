package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type cache struct {
	store map[int]int
	index int
	mu    sync.RWMutex
}

func (c *cache) addValue(value int) {
	findIdx := c.findIndex(value)
	c.mu.Lock()
	defer c.mu.Unlock()
	if findIdx == 0 {
		c.index++
		c.store[c.index] = value
	}
}

func (c *cache) findIndex(value int) int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	for k, v := range c.store {
		if v == value {
			return k
		}
	}
	return 0
}

func main() {

	cache := &cache{store: map[int]int{}}

	for i := 0; i < 10000; i++ {
		if i%5 == 0 {
			go func() {
				cache.addValue(rand.Intn(1000))

			}()
		} else {
			go func() {
				v := rand.Intn(1000)
				ind := cache.findIndex(v)
				if ind != 0 {
					fmt.Printf("element in the cache: index-%d  value-%d\n", ind, v)
				}
			}()
		}

	}

	fmt.Println(cache.store)
}
