package main

import (
	"fmt"
)

// checkSliceIsSorted проверяет, отсортирован ли слайс a по возрастанию.
func checkSliceIsSorted(a []int) bool {
	// Проходим по всем элементам слайса, кроме последнего
	for i := 0; i < len(a)-1; i++ {
		// Если текущий элемент больше следующего, слайс не отсортирован
		if a[i] > a[i+1] {
			return false
		}
	}
	// Если ни одно из условий не выполнено, слайс отсортирован
	return true
}

func main() {
	// Примеры использования функции
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 3, 4, 1, 2}
	slice3 := []int{1, 2, 2, 3, 4}

	fmt.Printf("Слайс %v отсортирован: %v\n", slice1, checkSliceIsSorted(slice1)) // true
	fmt.Printf("Слайс %v отсортирован: %v\n", slice2, checkSliceIsSorted(slice2)) // false
	fmt.Printf("Слайс %v отсортирован: %v\n", slice3, checkSliceIsSorted(slice3)) // true
}
