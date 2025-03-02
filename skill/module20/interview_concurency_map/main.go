package main

import (
	"fmt"
	"sync"
)

var (
	allData = make(map[string]string)
	rwm     sync.RWMutex
)

func Get(key string) string {
	rwm.RLock()
	defer rwm.RUnlock()
	return allData[key]
}
func Set(key string, value string) {
	rwm.Lock()
	defer rwm.Unlock()
	allData[key] = value
}
func main() {
	Set("a", "Hello")
	s := Get("a")
	fmt.Println(s)

}
