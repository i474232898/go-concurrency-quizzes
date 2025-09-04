# 🔗 Go Coding Task: Merging Two Sorted Channels

## Task Description

Implement a function `mergeSorted` that merges **two sorted input channels** into a single **sorted output channel**.

---

## 🔀 `mergeSorted(a, b <-chan int) <-chan int`

- Accepts:
  - Two **read-only channels of integers** `a` and `b`, each delivering values in sorted order.
- Returns:
  - A **read-only channel of integers** that emits all values from `a` and `b` in sorted order.

---

## 📈 Example Usage

```go
package main

import "fmt"

func mergeSorted(a, b <-chan int) <-chan int {
	// your code here
}

func fillChanA(c chan int) {
	c <- 1
	c <- 2
	c <- 4
	close(c)
}

func fillChanB(c chan int) {
	c <- -1
	c <- 4
	c <- 5
	close(c)
}

func main() {
	a, b := make(chan int), make(chan int)
	go fillChanA(a)
	go fillChanB(b)
	c := mergeSorted(a, b)

	for val := range c {
		fmt.Printf("%d ", val)
	}
}
```
