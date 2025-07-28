# ⚙️ Go Coding Task: Controlled Parallel Execution with Error Aggregation

## Task Description

Implement a concurrency control mechanism via a `waiter` interface that:

1. **Runs functions concurrently**, up to a limit specified by `maxParallel`.
2. Aggregates errors from those function executions using [`errors.Join`](https://pkg.go.dev/errors#Join).
3. Implements this using a struct `waitGroup` that manages synchronization, parallelism, and error collection.

---

## 🔧 `waiter` Interface

```go
package main

import (
	"context"
	"errors"
)

type waiter interface {
	wait() error
	run(ctx context.Context, f func(ctx context.Context) error)
}
type waitGroup struct {
  // your code here
}

func (g *waitGroup) wait() error {
  // your code here
}
func (g *waitGroup) run(ctx context.Context, fn func(ctx context.Context) error) {
  // your code here
}
func newGroupWait(maxParallel int) waiter {
  // your code here
}
func main() {
	g := newGroupWait(2)
  
	ctx := context.Background()
	expErr1 := errors.New("got error 1")
	expErr2 := errors.New("got error 2")
	g.run(ctx, func(ctx context.Context) error {
		return nil
	})
	g.run(ctx, func(ctx context.Context) error {
		return expErr2
	})
	g.run(ctx, func(ctx context.Context) error {
		return expErr1
	})
	err := g.wait()
	if !errors.Is(err, expErr1) || !errors.Is(err, expErr2) {
		panic("wrong code")
	}
}
```
