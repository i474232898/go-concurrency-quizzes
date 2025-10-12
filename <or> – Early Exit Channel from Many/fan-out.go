package main

import (
	"fmt"
	"sync"
	"time"
)

var p = fmt.Println

func or(channels ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})

	Once := sync.Once{}
	for _, ch := range channels {
		go func(c <-chan interface{}) {
			select {
			case <-c:
				Once.Do(func() {
					close(out)
				})
			case <-out:
				return
			}
		}(ch)
	}

	return out
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),

		sig(1*time.Second), // this one triggers first
		sig(2*time.Second),

		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	p("done after %v\n", time.Since(start))
	time.Sleep(3 * time.Second)
}
