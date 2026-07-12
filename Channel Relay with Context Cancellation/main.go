package main

import (
	"context"
	"fmt"
)

var p = fmt.Println

func orDone(ctx context.Context, in <-chan interface{}) <-chan interface{} {
	out := make(chan interface{})

	go func() {
		defer close(out)
		for {
			select {
			case data, ok := <-in:
				if !ok {
					return
				}
				select {
				case out <- data:
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return out
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan interface{})

	// Closed after orDone has received the second value from ch.
	secondValueTaken := make(chan struct{})

	go func() {
		ch <- 1
		ch <- 2

		close(secondValueTaken)

		close(ch)
	}()

	out := orDone(ctx, ch)

	v := <-out
	println("received:", v.(int))

	<-secondValueTaken

	cancel()

	_, ok := <-out
	if ok {
		// goroutine leak
		panic("expected out channel to be closed")
	}

	println("orDone exited successfully")
}
