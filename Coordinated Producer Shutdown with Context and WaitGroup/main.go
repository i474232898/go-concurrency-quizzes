package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

const (
	produceCount = 3
	produceStop  = 10
)

var p = fmt.Println

var i int64 = 0

func produce(pipe chan<- int, ctx context.Context) {
	defer func() {
		time.Sleep(3 * time.Second)
		p("produce finished")
	}()

	for {
		n := int(atomic.AddInt64(&i, 1) - 1)
		select {
		case pipe <- n:
		case <-ctx.Done():
			return
		}
	}
}

func main() {
	pipe := make(chan int)
	wg := sync.WaitGroup{}

	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(produceCount)

	for range produceCount {
		go func() {
			defer wg.Done()
			produce(pipe, ctx)
		}()
	}

	for data := range pipe {
		p(data)
		if data == produceStop {
			cancel()
			wg.Wait()

			p("main finished")
			return
		}
	}
}
