package main

import (
	"fmt"
	"sync"
)

func worker(f func(int) int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		results <- f(job)
	}
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
	wg.Add(numWorkers)
	for w := 0; w < numWorkers; w++ {
		go func() {
			defer wg.Done()
			worker(multiplier, jobs, results)
		}()
	}
	// lightweight alternative to sync.WaitGroup when you’re just waiting for one goroutine to finish.
	done := make(chan struct{})
	go func() {
		for r := range results {
			fmt.Println(r)
		}
		close(done)
	}()
	
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	go func(){
		wg.Wait()
		close(results)
	}()

	<-done
}
