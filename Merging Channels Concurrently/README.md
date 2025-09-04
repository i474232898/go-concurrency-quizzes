# 🧩 Go Coding Task: Merging Channels Concurrently

## Task Description

Implement two functions: `fillChan` and `merge`.

---

### 🔧 `fillChan(n int) <-chan int`

- Accepts an integer `n`.
- Returns a **read-only channel of integers**.
- Sends integers from `0` to `n-1` into the channel.

---

### 🔀 `merge(cs ...<-chan int) <-chan int`

- Accepts a variadic slice of read-only integer channels `cs`.
- Returns a **read-only channel of integers**.
- Concurrently reads values from all input channels and sends them into the returned channel.

---

## 🧪 Example Usage

```go
package main

import (
	"fmt"
	"sync"
)

// merge - merges multiple input channels into a single output channel
func merge(cs ...<-chan int) <-chan int {
	// your code here
}

// fillChan - fills a channel with integers from 0 to n-1
func fillChan(n int) <-chan int {
	// your code here
}

func main() {
	a := fillChan(2) // [0, 1]
	b := fillChan(3) // [0, 1, 2]
	c := fillChan(4) // [0, 1, 2, 3]

	d := merge(a, b, c)

	for v := range d {
		fmt.Println(v)
	}
}
```
