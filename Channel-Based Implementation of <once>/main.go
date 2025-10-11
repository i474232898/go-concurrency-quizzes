package main

import (
	"fmt"
	"sync"
)

var p = fmt.Println

const goroutinesNumber = 10

type once struct {
	c chan struct{}
}

func new() *once {
	c := make(chan struct{}, 1)
	c <- struct{}{}
	close(c)
	return &once{c}
}

func (o *once) do(f func()) {
	if _, ok := <-o.c; ok {
		f()
	}
}

func funcToCall() {
	p("call")
}

func main() {
	wg := sync.WaitGroup{}
	so := new()

	wg.Add(goroutinesNumber)

	for i := 0; i < goroutinesNumber; i++ {
		go func(f func()) {
			defer wg.Done()
			so.do(f)
		}(funcToCall)
	}

	wg.Wait()
}
