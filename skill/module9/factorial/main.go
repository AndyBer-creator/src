package main

import "fmt"

func main() {
	fmt.Println(factorial(7))
}

// расчет факториала интерактивно
func factorial(x int64) int64 {
	res := int64(1)
	for i := int64(1); i <= x; i++ {
		res *= i
	}
	return res
}

//Расчет факторала через рекурсию
// func FactR(n int64) int64{
// 	if n < 2 {
// 		return 1
// 	}
// 	return n*FactR(n-1)
// }
