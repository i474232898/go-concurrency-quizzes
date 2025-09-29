# 🔄 Go Coding Task: Coordinated Producer Shutdown with Context and WaitGroup

## 📝 Task Description

Implement two functions: `produce` and `main`.

---

### 🔧 `produce(pipe chan<- int, ctx context.Context)`

- Continuously sends increasing integers (starting from 0) into the shared `pipe` channel.
- Terminates early when the context is canceled.
- On termination:
  - Sleeps for 3 seconds.
  - Prints:  
    ```
    produce finished
    ```

---

### 🚀 `main()`

- Creates a shared channel `pipe`.
- Launches `produceCount` concurrent `produce` goroutines.
- Continuously reads from `pipe`, printing each number.
- When the value `produceStop` is received:
  - Cancels all producers via context.
  - Waits for all producers to print `"produce finished"`.
  - Finally prints:  
    ```
    main finished
    ```

---

## 📈 Example Usage

```go
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const (
	produceCount = 3
	produceStop  = 10
)

func produce(pipe chan<- int) { // you allowed to change signature
	// your implementation here
}

func main() {
	// your implementation here
}
```

---

## Output

 - Last 4 lines should be: 
```go
produce finished
produce finished
produce finished
main finished
```
