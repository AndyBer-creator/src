//Напишите программу, выводящую первые 20 простых чисел.

package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	counter := 1
	start := time.Now()
	for i := 3; counter <= 20000; i++ {
		c := false
		for j := 2; j < int(math.Sqrt((float64(i)))); j++ {
			if i%j == 0 {
				c = true
			}
			if c {
				break
			}
		}
		if !c {
			//fmt.Println(i)   //раскомментить при малых числах
			counter++
		}
	}
	end := time.Now()
	fmt.Printf("Finished: %s\n", end.Sub(start))
}
