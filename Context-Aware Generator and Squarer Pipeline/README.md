# 🔄 Go Coding Task: Context-Aware Generator and Squarer Pipeline

## Task Description

Implement two context-aware functions: `generator` and `squarer`.

---

### 🧮 `generator(ctx context.Context, in ...int) <-chan int`

- Accepts:
  - A `context.Context` to support cancellation.
  - A variadic list of integers.
- Returns a **read-only channel of integers**.
- Sends each value from the `in` slice into the channel in order.
- Terminates early if the context is canceled.

---

### 🔢 `squarer(ctx context.Context, in <-chan int) <-chan int`

- Accepts:
  - A `context.Context` to support cancellation.
  - A **read-only channel of integers**.
- Returns a **read-only channel of integers**.
- For each value received from the input channel:
  - Computes its square.
  - Sends the result into the output channel.
- Terminates early if the context is canceled.

---

## 🔧 Example Usage

```go
package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	pipeline := squarer(ctx, generator(ctx, 1, 2, 3))

	for x := range pipeline {
		fmt.Println(x)
	}
}

func generator(ctx context.Context, in ...int) <-chan int {
	// your code here
}

func squarer(ctx context.Context, in <-chan int) <-chan int {
	// your code here
}
