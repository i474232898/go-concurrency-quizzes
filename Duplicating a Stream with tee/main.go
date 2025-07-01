package main

import (
	"context"
	"reflect"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	i := 0
	inc := func() interface{} {
		i++
		return i
	}
	out1, out2 := tee(ctx, take(ctx, repeatFn(ctx, inc), 3))

	var res1, res2 []interface{}
	for val1 := range out1 {
		res1 = append(res1, val1)
		res2 = append(res2, <-out2)
	}

	exp := []interface{}{1, 2, 3}
	if !reflect.DeepEqual(res1, exp) || !reflect.DeepEqual(res2, exp) {
		panic("wrong code")
	}
}

// tee duplicates each value from `in` into two output channels.
func tee(ctx context.Context, in <-chan interface{}) (<-chan interface{}, <-chan interface{}) {
	outA, outB := make(chan interface{}), make(chan interface{})

	go func() {
		defer func() {
			close(outA)
			close(outB)
		}()

		for v := range orDone(ctx, in) {
			select {
			case outA <- v:
			case <-ctx.Done():
				return
			}
			select {
			case outB <- v:
			case <-ctx.Done():
				return
			}
		}
	}()

	return outA, outB
}

// orDone forwards values from `in` until `in` is closed or context is done.
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

// repeatFn repeatedly calls `fn` and sends the values until context is done.
func repeatFn(ctx context.Context, fn func() interface{}) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case out <- fn():
			}
		}
	}()
	return out
}

// take reads up to `num` values from `in` or stops if context is done.
func take(ctx context.Context, in <-chan interface{}, num int) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		defer close(out)
		for i := 0; i < num; i++ {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-in:
				if !ok {
					return
				}
				out <- v
			}
		}
	}()
	return out
}
