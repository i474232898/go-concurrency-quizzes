# 🏗️ Go Coding Task: Worker Pool with Goroutines and Channels

## Task Description

Implement a simple **worker pool** to run `numJobs` tasks in parallel, using `numWorkers` goroutines which are launched **only once** during the program’s execution.

---

## 🔨 `worker(f func(int) int, jobs <-chan int, results chan<- int)`

- Accepts:
  - A function `f` to execute.
  - A **read-only channel `jobs`** for receiving input values.
  - A **write-only channel `results`** for sending computed results.
- Continuously reads jobs from the `jobs` channel, computes `f(job)`, and writes the result into the `results` channel.

---

## 🚀 `main` Function

- Creates:
  - A `jobs` channel and a `results` channel, each with buffer size `numJobs`.
- Launches exactly `numWorkers` goroutines running the `worker` function, using `multiplier` as the function to apply.
- Sends integers from `1` to `numJobs` into the `jobs` channel.
- Reads and prints values from the `results` channel **as they are produced**, concurrently with the workers.

---

## 🔧 Example Stub

```go
package main

import (
	"fmt"
	"sync"
)

func worker(f func(int) int, jobs <-chan int, results chan<- int) {
	// your code here
}

const numJobs = 5
const numWorkers = 3

func main() {
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	wg := sync.WaitGroup{}

	multiplier := func(x int) int {
		return x * 10
	}

	// your code here
}
