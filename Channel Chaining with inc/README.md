# ➕ Go Coding Task: Channel Chaining with `inc`

## Task Description

Implement a function `inc` that:

- Accepts a read-only channel of integers.
- Returns a new channel that emits values from the input channel, incremented by 1.

Extend the `main` function to create a **chain of `n` `inc` functions**, each wrapping the previous channel. The final result should be:

- `0` → pass through `n` `inc` steps → final value should be `n`.

---

## 🔧 Function Signature

```go
package main

func main() {
	first := make(chan int)
	last := make(<-chan int)

	n := 10

  // your code here

	first <- 0
	close(first)

	if n != <-last {
		panic("wrong code")
	}
}

func inc(in <-chan int) <-chan int {

}
```

