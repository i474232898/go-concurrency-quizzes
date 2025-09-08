# ⏱️ Go Coding Task: Running a Task with Context Timeout

## Task Description

You are given a function `executeTask` that might hang indefinitely.

Your goal is to implement a **wrapper function** `executeTaskWithTimeout` that:

- Accepts a `context.Context`.
- Executes `executeTask`.
- Terminates either:
  - When `executeTask` finishes, returning `nil`.
  - Or when the context is done (e.g., due to timeout or cancellation), returning the context’s error.

---

## 🔧 `executeTaskWithTimeout(ctx context.Context) error`

- Should start `executeTask` in a separate goroutine.
- Should monitor the context for cancellation using `ctx.Done()`.
- If the context finishes before `executeTask` does, it should return `ctx.Err()`.

---

## 🧩 Example Usage

```go
package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

const timeout = 100 * time.Millisecond

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	err := executeTaskWithTimeout(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("task done")
}

func executeTaskWithTimeout(ctx context.Context) error {
	// your code here
}

func executeTask() {
	time.Sleep(time.Duration(rand.Intn(3)) * timeout)
}
```
