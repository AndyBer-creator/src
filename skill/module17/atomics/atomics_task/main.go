// Программа имеет условия возникновения гонок за данными, не гарантирующие вывод в консоль конечного значения счётчика (то есть 1000).
// Программа имеет условия возникновения гонок за данными, которые могут привести к deadlock.
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

const step int64 = 1
const interationAmount int = 1000

func main() {
	var counter int64 = 0
	var c = sync.NewCond(&sync.Mutex{})
	increment := func(i int) {
		atomic.AddInt64(&counter, step)
		if i == interationAmount {
			c.Signal()
		}
	}
	for i := 1; i <= interationAmount; i++ {
		go increment(i)
	}
	c.L.Lock()
	c.Wait()
	c.L.Unlock()
	fmt.Println(counter)
}
