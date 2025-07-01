package main

import "fmt"

func mergeSorted(a, b <-chan int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		valA, okA := <-a
		valB, okB := <-b

		for okA && okB {
			if valA < valB {
				ch <- valA
				valA, okA = <-a
			} else {
				ch <- valB
				valB, okB = <-b
			}
		}
		for okA {
			ch <- valA
			valA, okA = <-a
		}
		for okB {
			ch <- valB
			valB, okB = <-b
		}
	}()

	return ch
}

func fillChanA(c chan int) {
	c <- 1
	c <- 2
	c <- 4
	close(c)
}

func fillChanB(c chan int) {
	c <- -1
	c <- 4
	c <- 5
	close(c)
}

func main() {
	a, b := make(chan int), make(chan int)
	go fillChanA(a)
	go fillChanB(b)
	c := mergeSorted(a, b)

	for val := range c {
		fmt.Printf("%d ", val)
	}
}
