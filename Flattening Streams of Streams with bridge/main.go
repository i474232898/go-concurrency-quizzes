package main

import (
	"context"
	"fmt"
	"reflect"
)

func main() {
	genVals := func() <-chan <-chan interface{} {
		out := make(chan (<-chan interface{}))
		go func() {
			defer close(out)
			for i := 0; i < 3; i++ {
				stream := make(chan interface{}, 1)
				stream <- i
				close(stream)
				out <- stream
			}
		}()
		return out
	}

	var res []interface{}
	for v := range bridge(context.Background(), genVals()) {
		res = append(res, v)
		fmt.Println(res)
	}

	if !reflect.DeepEqual(res, []interface{}{0, 1, 2}) {
		panic("wrong code")
	}
}

// bridge reads channels from `ins`, and for each one forwards its values.
func bridge(ctx context.Context, ins <-chan <-chan interface{}) <-chan interface{} {
	out := make(chan any)

	go func() {
		defer close(out)
		for {
			var ch <-chan any
			select {
			case c, ok := <-ins:
				if !ok {
					return
				}
				ch = c
			case <-ctx.Done():
			}

			for d := range orDone(ctx, ch) {
				out <- d
			}
		}
	}()

	return out
}

// orDone reads from `in` until it is closed or the context is done.
func orDone(ctx context.Context, in <-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-in:
				if !ok {
					return
				}
				select {
				case out <- v:
				case <-ctx.Done():
				}
			}
		}
	}()
	return out
}
