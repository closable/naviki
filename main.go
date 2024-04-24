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
	// variant_1(20)
	// variant_2(7)

	variant_3()

}

func variant_3() {

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

func addChannel(done chan struct{}) chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < 10; i++ {
			select {
			case <-done:
				return
			case ch <- rand.Intn(10):
			}
		}
	}()
	return ch
}

func mergeChans(done chan struct{}, ch ...chan int) chan int {
	var wg = sync.WaitGroup{}
	res := make(chan int)

	for _, channel := range ch {
		//_ch := channel
		wg.Add(1)
		go func(chnl chan int) {
			defer wg.Done()
			for v := range chnl {
				select {
				case <-done:
					return
				case res <- v:
				}
			}
		}(channel)
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	return res
}

func variant_2(elements int) {
	count := elements
	done := make(chan struct{})
	chans := make([]chan int, count)

	for i := 0; i < count; i++ {
		ch := addChannel(done)
		chans[i] = ch
	}

	result := mergeChans(done, chans...)

	for v := range result {
		fmt.Printf("%d ", v)
	}
}

func variant_1(elements int) {
	fmt.Println("Вариант 1")
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		for i := 0; i < elements; i++ {
			v := rand.Intn(10)
			ch1 <- v
			t := rand.Intn(10)
			ch2 <- t
		}
		close(ch1)
		close(ch2)
	}()

	go func() {
	loop:
		for {
			select {
			case x, ok := <-ch1:
				if !ok {
					break loop
				}
				ch3 <- x
			case x, ok := <-ch2:
				if !ok {
					break loop
				}
				ch3 <- x
			}
		}
		close(ch3)
	}()

	for i := range ch3 {
		fmt.Printf("%d ", i)
	}
}
