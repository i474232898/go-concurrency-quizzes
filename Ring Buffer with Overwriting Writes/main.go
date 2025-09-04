package main

import (
	"fmt"
	"reflect"
)

var p = fmt.Println

type ringBuffer struct {
	c chan int
}

func newRingBuffer(size int) *ringBuffer {
	return &ringBuffer{c: make(chan int, size)}
}

func (b *ringBuffer) write(v int) {
	for {
		select {
		case b.c <- v:
			return
		default:
			<-b.c
		}
	}
}

func (b *ringBuffer) close() {
	close(b.c)
}

func (b *ringBuffer) read() (v int, ok bool) {
	select {
	case v, ok = <-b.c:
		return v, ok
	default:
		return 0, false
	}
}

func main() {
	buff := newRingBuffer(3)

	for i := 1; i <= 6; i++ {
		buff.write(i)
	}

	buff.close()

	res := make([]int, 0)
	for {
		if v, ok := buff.read(); ok {
			res = append(res, v)
		} else {
			break
		}
	}

	if !reflect.DeepEqual(res, []int{4, 5, 6}) {
		panic(fmt.Sprintf("wrong code, res is %v", res))
	}
}
