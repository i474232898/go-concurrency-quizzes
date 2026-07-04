# 🏗️ Go Coding Task: Worker Pool with Goroutines and Channels

## Task Description

Implement a simple **worker pool** to run `numJobs` tasks in parallel, using `numWorkers` goroutines which are launched **only once** during the program's execution.

Data flows in one direction:

```
main → jobs channel → [workers] → results channel → main
```

---

## 🔨 `worker(f func(int) int, jobs <-chan int, results chan<- int)`

- Accepts:
  - A function `f` to execute.
  - A **read-only channel `jobs`** for receiving input values.
  - A **write-only channel `results`** for sending computed results.
- Continuously reads jobs from the `jobs` channel, computes `f(job)`, and writes the result into the `results` channel.

---

## 🚀 `main` Function

- Creates a `jobs` channel and a `results` channel, each with buffer size `numJobs`.
- Launches exactly `numWorkers` goroutines running the `worker` function, using `multiplier` as the function to apply.
- Sends integers from `1` to `numJobs` into the `jobs` channel.
- Prints each result as it arrives. The main goroutine must block until all results are received and printed — no busy-waiting or sleeps.

---

## 🤔 Things to Consider

- How does a worker know there are no more jobs and it should stop?
- How does `main` know all results have been sent before it exits?
- Is the output order guaranteed?

---

## 📤 Expected Output

```
// possible output (order may vary):
10
30
20
40
50
```

---

## 🔧 Example Stub

```go
package main

import (
	"fmt"
)
var p = fmt.Println

func worker(f func(int) int, jobs <-chan int, results chan<- int) {
	// your code here
}

const numJobs = 5
const numWorkers = 3

func main() {
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	multiplier := func(x int) int {
		return x * 10
	}

	// your code here
}
```
