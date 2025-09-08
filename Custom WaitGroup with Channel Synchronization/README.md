# ⏳ Go Coding Task: Custom WaitGroup Using Channels

## 📝 Task Description

Implement a custom `Group` structure (similar to `sync.WaitGroup`) that:

- Tracks the number of goroutines.
- Provides:
  - `Done()` to mark a task as completed.
  - `Wait()` to wait for all tasks to finish.
- Uses **channels** for synchronization.
- Ensures that the `main()` function completes **without panic**.

---

## 🔧 Starter Code

```go
package main

import (
	"reflect"
	"sort"
	"sync"
)

type Group struct {
	c    chan struct{}
	size int
}

func New(size int) *Group {
	// your code here
}

func (s *Group) Done() {
	// your code here
}

func (s *Group) Wait() {
	// your code here
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	n := len(numbers)

	var res []int
	var mu sync.Mutex
	group := New(n)

	for _, num := range numbers {
		go func(num int) {
			defer group.Done()
			mu.Lock()
			res = append(res, num)
			mu.Unlock()
		}(num)
	}

	group.Wait()
	sort.IntSlice(res).Sort()

	if !reflect.DeepEqual(res, numbers) {
		panic("wrong code")
	}
}
```
