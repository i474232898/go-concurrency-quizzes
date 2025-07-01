package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type result struct {
	msg string
	err error
}

type search func() *result
type replicas []search

func fakeSearch(kind string) search {
	return func() *result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return &result{
			msg: fmt.Sprintf("%q result", kind),
		}
	}
}

func getFirstResult(ctx context.Context, replicas replicas) *result {
	cns, cancel := context.WithCancel(ctx)
	defer cancel()
	resultCh := make(chan *result)

	for _, r := range replicas {
		go func(repl search) {
			select {
			case resultCh <- repl():
			case <-cns.Done():
			}
		}(r)
	}

	select {
	case res := <-resultCh:
		cancel()
		cns.Done()
		return res
	case <-ctx.Done():
		return &result{
			err: ctx.Err(),
		}
	}
}

func getResults(ctx context.Context, replicaKinds []replicas) []*result {
	res := make(chan *result)
	var wg sync.WaitGroup

	for _, rk := range replicaKinds {
		wg.Add(1)
		go func(repKind replicas) {
			res <- getFirstResult(ctx, repKind)
			wg.Done()
		}(rk)
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	arr := make([]*result, 0, len(replicaKinds))
	for r := range res {
		arr = append(arr, r)
	}
	return arr
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	replicaKinds := []replicas{
		{fakeSearch("web1"), fakeSearch("web2")},
		{fakeSearch("image1"), fakeSearch("image2")},
		{fakeSearch("video1"), fakeSearch("video2")},
	}

	for _, res := range getResults(ctx, replicaKinds) {
		fmt.Println(res.msg, res.err)
	}
}
