# 🔁 Go Coding Task: Ring Buffer with Overwriting Writes

## Task Description

Implement a `ringBuffer` structure with the following behavior:

- Acts like a circular buffer (ring buffer).
- Maintains **only the latest N values** (overwrites oldest when full).
- Provides `write`, `read`, and `close` methods.
- After closing, `read` should return remaining values until empty.

---

## 🔧 Required API

```go
package main

import (
	"fmt"
	"reflect"
)

type ringBuffer struct {
	c chan int
}

func newRingBuffer(size int) *ringBuffer {
	// your code here
	return nil
}

func (b *ringBuffer) write(v int) {
	// your code here
}

func (b *ringBuffer) close() {
	// your code here
}

func (b *ringBuffer) read() (v int, ok bool) {
	// your code here
	return 0, false
}

func main() {
	buff := newRingBuffer(3)

	for i := 1; i <= 6; i++ {
		buff.write(i)
	}

	buff.close()

	res := make([]int, 0)
	for {
		if v, ok := buff.read(); ok {
			res = append(res, v)
		} else {
			break
		}
	}

	if !reflect.DeepEqual(res, []int{4, 5, 6}) {
		panic(fmt.Sprintf("wrong code, res is %v", res))
	}
}
```
