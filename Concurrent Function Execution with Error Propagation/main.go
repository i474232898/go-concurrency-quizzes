package main

import (
	"errors"
	"fmt"
	"sync"
)

var p = fmt.Println

type fn func() error

func Run(fs ...fn) error {
	var wg sync.WaitGroup
	wg.Add(len(fs))
	var errr error = nil
	var once sync.Once

	for _, fnc := range fs {
		go func(fnn fn) {
			defer wg.Done()
			if err := fnn(); err != nil {
				once.Do(func() {
					errr = err
				})
			}
		}(fnc)
	}
	wg.Wait()

	return errr
}

func main() {
	expErr := errors.New("error")

	funcs := []fn{
		func() error { return nil },
		func() error { return nil },
		func() error { return expErr },
		func() error { return nil },
	}

	if err := Run(funcs...); !errors.Is(err, expErr) {
		panic("wrong code")
	}
}
