package main

import (
	"context"
	"fmt"
	"time"
)

func someFunc(ctx context.Context) {
	for {
		fmt.Println(time.Now())
		time.Sleep(time.Millisecond * 100)
		fmt.Println("start some func")
		select {
		case <-ctx.Done():
			return
		default:
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go someFunc(ctx)

	fmt.Println("start wait")
	time.AfterFunc(time.Second, cancel)
	fmt.Println("finish wait")
	time.Sleep(time.Second * 2)
}
