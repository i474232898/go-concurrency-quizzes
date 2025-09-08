package main

import (
	"fmt"
	"reflect"
	"sort"
	"sync"
)

var p = fmt.Println

type Group struct {
	c    chan struct{}
	size int
}

func New(size int) *Group {
	return &Group{make(chan struct{}, size), size}
}

func (s *Group) Done() {
	s.c <- struct{}{}
}

func (s *Group) Wait() {
	for i := 0; i < s.size; i++ {
		<-s.c
	}
	close(s.c)
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	n := len(numbers)

	var res []int
	var mu sync.Mutex
	group := New(n)

	for _, num := range numbers {
		go func(num int) {
			defer group.Done()
			mu.Lock()
			res = append(res, num)
			mu.Unlock()
		}(num)
	}

	group.Wait()
	sort.IntSlice(res).Sort()

	if !reflect.DeepEqual(res, numbers) {
		panic("wrong code")
	}
}
