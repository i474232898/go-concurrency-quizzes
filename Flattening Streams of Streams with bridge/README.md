# 🌉 Go Coding Task: Flattening Streams of Streams with `bridge`

## Task Description

Implement a function `bridge` that:

- Receives a **channel of channels** (i.e., a stream of streams) called `ins`.
- Continuously reads each inner channel from `ins`.
- For each inner channel, reads all its values and forwards them into a single flattened output channel.
- Stops early if the context is canceled.

---

## 🔧 `bridge(ctx context.Context, ins <-chan <-chan interface{}) <-chan interface{}`

- Accepts:
  - A `context.Context` for cancellation.
  - A channel `ins` that yields **channels of `interface{}`**.
- Returns:
  - A **single output channel** that sequentially emits values from each inner channel.

---

## 📈 Example Usage

```go
package main

import (
	"context"
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
	// your code here
	return nil
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
