package main

import (
	"context"
	"fmt"
	"math/rand"
)

func repeatFn(ctx context.Context, fn func() interface{}) <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		for {
			select {
			case <-ctx.Done():
				return
			case ch <- fn():
			}
		}
	}()

	return ch
}

func take(ctx context.Context, in <-chan interface{}, num int) <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		for i := 0; i < num; i++ {
			select {
			case <-ctx.Done():
				return
			case v, ok := <- in:
				if !ok {
					return 
				}
				ch <- v
			}
		}
	}()

	return ch
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	randFn := func() interface{} { return rand.Int() }
	var res []interface{}

	for num := range take(ctx, repeatFn(ctx, randFn), 3) {
		res = append(res, num)
	}

	if len(res) != 3 {
		panic("wrong code")
	}

	fmt.Println("Generated values:", res)
}
