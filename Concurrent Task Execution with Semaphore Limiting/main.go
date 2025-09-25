package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var p = fmt.Println

func processTask() {
	time.Sleep(100 * time.Millisecond) // long job simulation
}

func main() {
	// cpu := runtime.NumCPU()
	cpu := 3 // or use custom limit
	ch := make(chan struct{}, cpu)
	wg := sync.WaitGroup{}

	producerDone := make(chan struct{})
	go func() {
		for i := 0; i < 50; i++ {
			ch <- struct{}{}
			wg.Add(1)
			go func() {
				defer func() {
					<-ch
					wg.Done()
				}()
				processTask() // do it in parallel limiting by CPU number
			}()
		}
		close(producerDone)
	}()
	<-producerDone

	p("goroutines number: ", runtime.NumGoroutine())
	wg.Wait()
}
