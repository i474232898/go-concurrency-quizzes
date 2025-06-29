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

// fakeDownload simulates downloading from a URL.
func fakeDownload(url string) error {
	// Simulate a network operation...
	return nil
}

func download(urls []string) error {
	// your code here
	return nil
}
