package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

func main() {
	ar := make([]int, 50)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайное значение от [-100;100]
	}

	bubbleSort(ar)
	fmt.Printf("Sorting by ascending...:\t%v\n", ar)
	// bubbleSortReversed(ar)
	// fmt.Printf("Sorting by descending ...:\t%v\n", ar)
	// bubbleSortRecursive(ar)
	// fmt.Printf("Sorting with recursion...:\t%v\n", ar)
}

func bubbleSort(ar []int) {
	// ваш код здесь
	fmt.Printf("Unsorted list:\t%v\n", ar)
	fmt.Println("")
	length := len(ar)
	for i := 0; i < (length - 1); i++ {
		for j := 0; j < ((length - 1) - i); j++ {
			if ar[j] > ar[j+1] {
				ar[j], ar[j+1] = ar[j+1], ar[j]
			}
		}

	}

}

// func bubbleSortReversed(ar []int) {
// 	length := len(ar)
// 	for i := 0; i < (length - 1); i++ {
// 		for j := 0; j < ((length - 1) - i); j++ {
// 			if ar[j] < ar[j+1] {
// 				ar[j], ar[j+1] = ar[j+1], ar[j]
// 			}
// 		}

// 	}
// }
// func bubbleSortRecursive(ar []int) {
// 	n := len(ar)
// 	if n <= 1 {
// 		return
// 	}

// 	// Выполняем один проход по массиву и перемещаем наибольший элемент в конец
// 	for i := 0; i < n-1; i++ {
// 		if ar[i] > ar[i+1] {
// 			// Меняем местами элементы
// 			ar[i], ar[i+1] = ar[i+1], ar[i]
// 		}
// 	}

// 	// Рекурсивно вызываем функцию для оставшейся части массива
// 	bubbleSortRecursive(ar[:n-1])

// }
