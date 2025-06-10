package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

var p = fmt.Println

type Result struct {
	msg string
	err error
}

func fakeDownload(url string) Result {
	n := rand.IntN(100)
	wait := time.Duration(n * int(time.Microsecond))
	time.Sleep(wait)

	if n > 50 {
		return Result{
			err: errors.New(fmt.Sprintf("url %s", url)),
		}
	}

	return Result{
		msg: fmt.Sprintf("downloaded %s", url),
	}
}

func download(urls []string) ([]string, error) {
	var wg sync.WaitGroup
	resultCh := make(chan Result, len(urls))
	wg.Add(len(urls))

	for _, u := range urls {
		go func(url string) {
			defer wg.Done()
			resultCh <- fakeDownload(url)
		}(u)
	}
	wg.Wait()
	close(resultCh)

	var messages []string
	var errs []error
	for res := range resultCh {
		if res.err != nil {
			errs = append(errs, res.err)
		} else {
			messages = append(messages, res.msg)
		}
	}
	err := errors.Join(errs...)

	return messages, err
}

func main() {
	result, _ := download([]string{
		"https://example.com/1.xml",
		"https://example.com/5.xml",
		"https://example.com/6.xml",
		"https://example.com/7.xml",
		"https://example.com/8.xml",
	})

	p(result)
}
