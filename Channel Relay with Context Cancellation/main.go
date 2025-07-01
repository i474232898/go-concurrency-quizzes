package main

import (
	"context"
	"reflect"
)

func orDone(ctx context.Context, in <-chan interface{}) <-chan interface{} {
	ch := make(chan interface{})

	go func() {
		defer close(ch)
		for {
			select {
			case data, ok := <-in:
				if !ok {
					return
				}
				select {
				case ch <- data:
				case <-ctx.Done():
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return ch
}

func main() {
	ch := make(chan interface{})
	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}
		close(ch)
	}()

	var res []interface{}
	for v := range orDone(context.Background(), ch) {
		res = append(res, v)
	}

	if !reflect.DeepEqual(res, []interface{}{0, 1, 2}) {
		panic("wrong code")
	}
}
