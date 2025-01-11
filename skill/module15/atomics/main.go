package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var v atomic.Value

type Config struct {
	a []int
}

func main() {
	// cfg := &Config{}
	// lock := &sync.Mutex{}

	go func() {
		var i int
		for {
			i++
			cfg := &Config{
				a: []int{i, i + 1, i + 2, i + 3, i + 4, i + 5},
			}
			v.Store(cfg)
		}
	}()

	var wg sync.WaitGroup

	for n := 0; n < 4; n++ {
		wg.Add(1)
		go func() {
			for n := 0; n < 100; n++ {
				cfg := v.Load()
				fmt.Println(cfg)
			}
			wg.Done()
		}()
	}

	wg.Wait()
}
