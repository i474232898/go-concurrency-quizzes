# 📚 Go Coding Task: Concurrent Sorted Head from Multiple Readers

## Task Description

Implement a function `ConcurrentSortHead` that:

- Accepts multiple `io.Reader`s, each providing **ascending sorted strings** (line by line).
- Reads from all readers **concurrently**.
- Returns the first `m` lines in **global ascending order**.

---

## 🔧 Function Signature

```go
package main

import (
	"io"
	"reflect"
	"strings"
)

func main() {
	f1 := `aaa
ddd
`
	f2 := `bbb
eee
`
	f3 := `ccc
fff
`

	files := []io.Reader{
		strings.NewReader(f1),
		strings.NewReader(f2),
		strings.NewReader(f3),
	}
	rows, err := ConcurrentSortHead(4, files...)
	if err != nil {
		panic(err)
	}
	if !reflect.DeepEqual(rows, []string{"aaa", "bbb", "ccc", "ddd"}) {
		panic("wrong code")
	}
}

func ConcurrentSortHead(m int, files ...io.Reader) ([]string, error) {
	// your code here
}
```
