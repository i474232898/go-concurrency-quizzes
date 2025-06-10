package main

import (
	"fmt"
	"sync"
	"time"
)

var p = fmt.Println

func fillChan(n int) <-chan int {
	ch := make(chan int)

	go func() {
		for i := range n {
			ch <- i
		}
		close(ch)
	}()

	return ch
}

func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	collect := func(ch <-chan int) {
		for n := range ch {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))

	for _, c := range cs {
		go collect(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	a := fillChan(2) // [0, 1]
	b := fillChan(3) // [0, 1, 2]
	c := fillChan(4) // [0, 1, 2, 3]

	d := merge(a, b, c)
	for v := range d {
		p(v)
	}
	time.Sleep(500 * time.Millisecond)
}
