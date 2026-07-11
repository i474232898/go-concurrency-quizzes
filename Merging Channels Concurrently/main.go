package main

import (
	"fmt"
	"sync"
	"time"
)

// merge - merges multiple input channels into a single output channel
func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)
	wg := sync.WaitGroup{}

	for _, ch := range cs {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for d := range c {
				out <- d
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// fillChan - fills a channel with integers from 0 to n-1
func fillChan(n int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := range n {
			out <- i
		}
	}()
	return out
}

func main() {
	a := fillChan(2) // [0, 1]
	b := fillChan(3) // [0, 1, 2]
	c := fillChan(4) // [0, 1, 2, 3]

	d := merge(a, b, c)

	for v := range d {
		fmt.Println(v)
	}

	time.Sleep(500 * time.Millisecond)
}
