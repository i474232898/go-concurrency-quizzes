# 📦 Go Coding Task: Concurrent Downloads with Error Aggregation

## Task Description

Implement a function called `download`.

---

### 🌐 `download(urls []string) error`

- Accepts a slice of URLs to download from.
- Concurrently downloads data from **each URL** using `fakeDownload` function.
- If any of the `fakeDownload` calls return an error, collect **all errors** and return them using [`errors.Join`](https://pkg.go.dev/errors#Join) (available since Go 1.20).
- If all downloads succeed, return `nil`.

---

## 🔧 Example Stub

```go
package main

import (
	"errors"
)

type Result struct {
	msg string
	err error
}

// fakeDownload simulates downloading from a URL.
func fakeDownload(url string) Result {
	n := rand.IntN(100)
	wait := time.Duration(n * int(time.Microsecond))
	time.Sleep(wait)

	if n > 50 {
		return Result{
			err: errors.New(fmt.Sprintf("url %s", url)),
		}
	}

	return Result{
		msg: fmt.Sprintf("downloaded %s", url),
	}
}

func download(urls []string) error {
	// your code here
	return nil
}
```
