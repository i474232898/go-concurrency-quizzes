package main

import (
	"context"
	"fmt"
)

var p = fmt.Println

func generator(ctx context.Context, in ...int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for _, v := range in {
			select {
			case ch <- v:
			case <-ctx.Done():
				return
			}
		}
	}()

	return ch
}

func squarer(ctx context.Context, in <-chan int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for {
			select {
			case v, ok := <-in:
				if !ok {
					return
				}
				ch <- v * v
			case <-ctx.Done():
				return
			}
		}
	}()

	return ch
}

func main() {
	ctx, _ := context.WithCancel(context.Background())

	pipeline := squarer(ctx, generator(ctx, 1, 2, 3))

	for i := range pipeline {
		p(i)
	}
}
