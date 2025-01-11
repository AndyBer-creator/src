package main

import (
	"fmt"
	"sync"
)

var wg = &sync.WaitGroup{}

func someFunc() {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(t int) {
			defer wg.Done()
			fmt.Println(t)
		}(i)
	}
}

func main() {
	// fmt.Println(1)
	someFunc()
	// time.Sleep(time.Second)
	wg.Wait()
}
