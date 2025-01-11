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
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайно значение от [-100;100]
	}

	ar = mergeSort(ar)

	fmt.Println(ar)
}

func mergeSort(ar []int) []int {
	// ваш код здесь
	if len(ar) <= 1 {
		return ar // Массив уже отсортирован
	}

	// Делим массив на две половины
	mid := len(ar) / 2
	left := mergeSort(ar[:mid])
	right := mergeSort(ar[mid:])

	// Объединяем отсортированные подмассивы
	return merge(left, right)
}

// merge объединяет два отсортированных подмассива в один отсортированный массив
func merge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0

	// Сравниваем элементы из обоих подмассивов и добавляем меньший в результат
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Добавляем оставшиеся элементы из левого подмассива
	for i < len(left) {
		result = append(result, left[i])
		i++
	}

	// Добавляем оставшиеся элементы из правого подмассива
	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	return result
}
