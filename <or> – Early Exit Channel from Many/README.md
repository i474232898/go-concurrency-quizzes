# 🧪 Go Coding Task: `or` – Early Exit Channel from Many

## 📝 Task Description

Implement the function `or` that combines multiple channels into one:

- It should return a single channel that:
  - Closes as soon as **any** of the provided `channels`:
    - Becomes readable (i.e., receives a value), or
    - Is closed.

> This is commonly known as the **"or-channel pattern"** in Go.

---

## 🔧 Function Signature

```go
package main

import (
	"fmt"
	"time"
)

var p = fmt.Println

func or(channels ...<-chan interface{}) <-chan interface{} {
	// implement here
	return nil
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
```
