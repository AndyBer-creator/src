// Напишите функцию, которая находит точку пересечения двух прямых,
//
//	заданных как y = ax + b на плоскости.
//	Коэффициенты a1, b1, a2, b2 задайте в начале программы
//	 или введите из консоли. Если прямые не пересекаются,
//	  вывести сообщение об этом.
package main

import (
	"fmt"
)

func main() {
	var x, y float64
	var a1, b1 float64 = 3, 4
	var a2, b2 float64 = 1, 4
	if a1 == a2 {
		fmt.Println("no crossing")
		return
	}
	x = (b2 - b1) / (a1 - a2)
	y = a1*x + b1
	fmt.Printf("x: %f, y: %f\n", x, y)
}
