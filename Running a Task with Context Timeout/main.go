package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

const timeout = 100 * time.Millisecond

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	err := executeTaskWithTimeout(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("task done")
}

func executeTaskWithTimeout(ctx context.Context) error {
	ch := make(chan struct{})
	go func ()  {
		executeTask()
		ch <- struct{}{}
	}()

	select {
	case <-ch:
		return  nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func executeTask() {
	time.Sleep(time.Duration(rand.Intn(3)) * timeout)
}
