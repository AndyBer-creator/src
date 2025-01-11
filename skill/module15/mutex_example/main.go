   package main

import (
	"fmt"
	"sync"
)

var wg = &sync.WaitGroup{}
var mutex = &sync.Mutex{}

func someFunc(i int64, m map[int64]int64) {
	defer wg.Done()
	mutex.Lock()
	defer mutex.Unlock()
	m[i] = i
}

func main() {
	var i int64
	m := &map[int64]int64{}
	for i = 0; i < 20; i++ {
		wg.Add(1)
		go someFunc(i, *m)
	}

	wg.Wait()
	fmt.Println(m)
}
