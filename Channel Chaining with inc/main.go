package main

func main() {
	first := make(chan int)
	last := make(<-chan int)

	n := 10

	last = inc(first)
	for i := 1; i < n; i++ {
		last = inc(last)
	}

	first <- 0
	close(first)

	if n != <-last {
		panic("wrong code")
	}
}

func inc(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in {
			out <- (v + 1)
		}
	}()

	return out
}
