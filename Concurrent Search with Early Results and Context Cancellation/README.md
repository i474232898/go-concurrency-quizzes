# 🚀 Go Coding Task: Concurrent Search with Early Results and Context Cancellation

## Task Description

You will implement two functions: `getFirstResult` and `getResults`.

---

## 🔍 `getFirstResult(ctx context.Context, replicas replicas) *result`

- Accepts:
  - A `context.Context` to support cancellation.
  - A slice of `searh` functions (named `replicas`), each simulating a replica performing a search.
- Concurrently starts all replica searches.
- Returns **the first available `*result`** from the replicas.
- If the `context` finishes before any search returns, returns a `*result` with `err` set to the context error.

---

## 🗂️ `getResults(ctx context.Context, replicaKinds []replicas) []*result`

- Accepts:
  - A `context.Context`.
  - A slice of `replicas` slices (i.e., multiple replica groups, each of type `replicas`).
- For each group in `replicaKinds`:
  - Starts a concurrent search using `getFirstResult`.
- Collects and returns a slice of results (one per replica group).

---

## 🧩 Example Usage

```go
package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type result struct {
	msg string
	err error
}

type searh func() *result
type replicas []searh

func fakeSearch(kind string) searh {
	return func() *result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return &result{
			msg: fmt.Sprintf("%q result", kind),
		}
	}
}

func getFirstResult(ctx context.Context, replicas replicas) *result {
	// your code here
}

func getResults(ctx context.Context, replicaKinds []replicas) []*result {
	// your code here
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	replicaKinds := []replicas{
		{fakeSearch("web1"), fakeSearch("web2")},
		{fakeSearch("image1"), fakeSearch("image2")},
		{fakeSearch("video1"), fakeSearch("video2")},
	}

	for _, res := range getResults(ctx, replicaKinds) {
		fmt.Println(res.msg, res.err)
	}
}
