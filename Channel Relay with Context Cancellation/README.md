# 🔄 Go Coding Task: Channel Relay with Context Cancellation (`orDone`)

## Task Description

Implement a function `orDone` that forwards values from a channel into a new channel, but stops early if the provided context is canceled.

---

## 🔧 `orDone(ctx context.Context, in <-chan interface{}) <-chan interface{}`

- Accepts:
  - A `context.Context` to support cancellation.
  - A **read-only input channel** `in` of type `interface{}`.
- Returns:
  - A **read-only output channel** that emits values received from `in`.
- Stops reading from `in` and closes the output channel when:
  - The input channel `in` is closed, **or**
  - The context is canceled (e.g., by timeout or manual cancel).

---

## 🧩 Example Usage

```go
package main

import (
	"context"
	"reflect"
)

func orDone(ctx context.Context, in <-chan interface{}) <-chan interface{} {
	// your code here
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
