# �� Go Coding Task: Concurrent Task Execution with Semaphore Limiting

## Task Description

Implement a **semaphore-based concurrency control** pattern that:

1. **Limits concurrent execution** to the number of CPU cores.

---

## 🔧 Example Stub

```go
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var p = fmt.Println

func processTask() {
	time.Sleep(100 * time.Millisecond) // long job simulation
}

func main() {
	// cpu := runtime.NumCPU()
	// cpu := 3  // or use custom limit

	for i := 0; i < 50; i++ {
		// processTask() // do it in parallel limiting by CPU number
	}

	p("goroutines number: ", runtime.NumGoroutine())
}
```

---

## �� Expected Behavior

- **Without concurrency**: 50 tasks × 100ms = ~5 seconds total time, 1 goroutine
- **With semaphore limiting**: Tasks execute concurrently but limited by CPU count, significantly faster execution, controlled goroutine count
