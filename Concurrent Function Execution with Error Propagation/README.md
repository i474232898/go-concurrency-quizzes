# 🚦 Go Coding Task: Concurrent Function Execution with Error Propagation

## Task Description

Implement the function `Run`, which executes a list of functions **concurrently** and waits for all of them to complete.

- If **any** function returns an error, `Run` should return **one of those errors** (any of them is acceptable).
- If **all functions succeed**, `Run` should return `nil`.

---

## 🔧 Function Signature

```go
package main

import (
	"errors"
)

type fn func() error

func Run(fs ...fn) error {
	// your code here
}

func main() {
	expErr := errors.New("error")

	funcs := []fn{
		func() error { return nil },
		func() error { return nil },
		func() error { return expErr },
		func() error { return nil },
	}

	if err := Run(funcs...); !errors.Is(err, expErr) {
		panic("wrong code")
	}
}
```
