package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var count int

	fmt.Println("Enter gorutines count")
	_, err := fmt.Scan(&count)
	checkerr(err) // Огромный минус, если студенты не обрабатывают ошибки. Ожидается что у студента везде будет if err != nil

	for i := 0; i < count; i++ {
		i := i
		go func() {
			for {
				fmt.Println(i)
				time.Sleep(time.Second)
			}
		}()
	}
	b := make([]byte, 1)
	os.Stdin.Read(b)
}

func checkerr(err error) {
	if err != nil {
		fmt.Println("FATAL: ", err.Error())
		os.Exit(1)
	}
}
