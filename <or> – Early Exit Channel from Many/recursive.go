package main

import (
	"fmt"
	"time"
)

var p = fmt.Println

func or(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	out := make(chan interface{})
	go func() {
		defer close(out)
		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-or(append(channels[2:], out)...):
			}
		}
	}()

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
