# 🔁 Go Coding Task: Channel-Based Implementation of `once`

## 📝 Task Description

Implement a `once` structure (similar to `sync.Once`) **using only channels**, without using the `sync` package (except in `main` for `WaitGroup`).

---

## 🔧 Requirements

### ✅ `type once`

A custom structure to ensure a function is executed **only once**, even when called by multiple goroutines.

### ✅ `func new() *once`

Returns a new pointer to a `once` instance. Internally, it must use **channels** for synchronization.

### ✅ `func (o *once) do(f func())`

- Accepts a function `f`.
- Ensures `f` is executed **only once**, no matter how many goroutines call `do`.
- All subsequent `do(f)` calls are no-ops.

---

## 🧩 Example Usage

```go
package main

import (
	"fmt"
	"sync"
)
var p = fmt.Println
const goroutinesNumber = 10

type once struct {
	// your implementation here (channel-based)
}

func new() *once {
	// your implementation here
	return nil
}

func (o *once) do(f func()) {
	// your implementation here
}

func funcToCall() {
	p("call")
}

func main() {
	wg := sync.WaitGroup{}
	so := new()

	wg.Add(goroutinesNumber)

	for i := 0; i < goroutinesNumber; i++ {
		go func(f func()) {
			defer wg.Done()
			so.do(f)
		}(funcToCall)
	}

	wg.Wait()
}
```
