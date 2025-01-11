package main

import (
	"fmt"
	"sync"
	"time"
)

var m sync.Map

func someFunc(i int64) {
	m.Store(i, i)

	val, ok := m.Load(i)
	if ok {
		fmt.Println(val)
	}
	// m.Delete(i)
}

func main() {
	var i int64
	for i = 0; i < 40; i++ {
		go someFunc(i)
	}

	time.Sleep(time.Second)

	m.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})
}
