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
	"fmt"
)

var p = fmt.Println

func orDone(ctx context.Context, in <-chan interface{}) <-chan interface{} {
	// your code here
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
```
