package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	variant_1(20)
	variant_2(7)
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
	fmt.Println("\nВариант 2")
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
