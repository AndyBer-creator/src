package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

type RingIntBuffer struct {
	array []int
	pos   int
	size  int
	m     sync.Mutex
}

func NewRingIntBuffer(size int) *RingIntBuffer {
	return &RingIntBuffer{make([]int, size), -1, size, sync.Mutex{}}
}
func (r *RingIntBuffer) Push(el int) {
	r.m.Lock()
	defer r.m.Unlock()
	if r.pos == r.size-1 {
		for i := 1; r.pos == r.size-1; i++ {
			r.array[i-1] = r.array[i]
		}
		r.array[r.pos] = el
	} else {
		r.pos++
		r.array[r.pos] = el
	}
}
func (r *RingIntBuffer) Get() []int {
	if r.pos <= 0 {
		return nil
	}
	r.m.Lock()
	defer r.m.Unlock()
	var output []int = r.array[:r.pos]

	r.pos = 0
	return output
}
func read(input chan<- int) {
	for {
		var u int
		_, err := fmt.Scanf("%d\n", &u)
		if err != nil {
			fmt.Println("не цифра")
		} else {
			input <- u
		}

	}
}
func removeNegative(currentChan <-chan int, nextChan chan<- int) {
	for number := range currentChan {
		if number >= 0 {
			nextChan <- number
		}
	}
}
func removeDivTree(currentChan <-chan int, nextChan chan<- int) {
	for number := range currentChan {
		if number%3 == 0 {
			nextChan <- number
		}

	}
}
func writeToBuffer(currentChan <-chan int, r *RingIntBuffer) {
	for number := range currentChan {
		r.Push(number)
	}
}

func witeToConsole(r *RingIntBuffer, t *time.Ticker) {
	for range t.C {
		buffer := r.Get()
		if len(buffer) > 0 {
			fmt.Println("Thebuffer: ", buffer)
		}
	}
}
func main() {
	input := make(chan int)
	go read(input)
	negativeFilterChan := make(chan int)
	go removeNegative(input, negativeFilterChan)
	divTreeChannel := make(chan int)
	go removeDivTree(negativeFilterChan, divTreeChannel)
	size := 20
	r := NewRingIntBuffer(size)
	go writeToBuffer(divTreeChannel, r)
	delay := 1
	ticker := time.NewTicker(time.Second * time.Duration(delay))
	go witeToConsole(r, ticker)
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	select {
	case sig := <-c:
		fmt.Printf("Got %s signal, Aborting ... \n", sig)
		os.Exit(0)
	}
}
