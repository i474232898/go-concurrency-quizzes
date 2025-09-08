# 🔁 Go Coding Task: repeatFn and take with Context Cancellation

## Task Description

Implement two functions: `repeatFn` and `take`.

---

### 🔄 `repeatFn(ctx context.Context, fn func() interface{}) <-chan interface{}`

- Accepts:
  - A `context.Context` to support cancellation.
  - A function `fn` that returns `interface{}`.
- Returns a **read-only channel of `interface{}`**.
- Continuously calls `fn()` in an infinite loop, sending each result into the returned channel.
- Stops early if the context is canceled.

---

### 🎯 `take(ctx context.Context, in <-chan interface{}, num int) <-chan interface{}`

- Accepts:
  - A `context.Context` to support cancellation.
  - A **read-only channel of `interface{}`**.
  - An integer `num` specifying the maximum number of values to read.
- Returns a **read-only channel of `interface{}`**.
- Reads up to `num` values from `in` and sends them into the returned channel.
- Stops early if the context is canceled.

---

## 🔧 Example Usage

```go
package main

import (
	"context"
	"fmt"
	"math/rand"
)

func repeatFn(ctx context.Context, fn func() interface{}) <-chan interface{} {
	// your code here
}

func take(ctx context.Context, in <-chan interface{}, num int) <-chan interface{} {
	// your code here
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
```
