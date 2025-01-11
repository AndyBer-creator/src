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
	arr := make([]int, 50)
	for i := range arr {
		arr[i] = rand.Intn(200) - 100 // ограничиваем случайное значение от [-100;100]
	}

	selectionSortByMin(arr)
	fmt.Println(arr)

	selectionSortByMax(arr)
	fmt.Println(arr)
	bidirectionalSelectionSort(arr)
	fmt.Println(arr)
}

func selectionSortByMin(arr []int) {
	// ваш код здесь
	n := len(arr)
	for i := 0; i < n-1; i++ {
		// Предполагаем, что текущий элемент является минимальным
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j // Находим индекс минимального элемента
			}
		}
		// Меняем местами текущий элемент с найденным минимальным
		if minIndex != i {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}
}
func selectionSortByMax(arr []int) {
	// ваш код здесь
	n := len(arr)
	for i := 0; i < n-1; i++ {
		// Предполагаем, что текущий элемент является максимальным
		maxIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] > arr[maxIndex] {
				maxIndex = j // Находим индекс максимального элемента
			}
		}
		// Меняем местами текущий элемент с найденным максимальным
		if maxIndex != i {
			arr[i], arr[maxIndex] = arr[maxIndex], arr[i]
		}
	}
}
func bidirectionalSelectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n/2; i++ {
		minIndex := i
		maxIndex := i

		// Находим минимальный и максимальный элементы
		for j := i; j < n-i; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j // Находим индекс минимального элемента
			}
			if arr[j] > arr[maxIndex] {
				maxIndex = j // Находим индекс максимального элемента
			}
		}

		// Меняем местами минимальный элемент с текущим
		if minIndex != i {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
			// Если максимальный элемент был на месте минимального, обновляем его индекс
			if maxIndex == i {
				maxIndex = minIndex
			}
		}

		// Меняем местами максимальный элемент с текущим
		if maxIndex != n-i-1 {
			arr[n-i-1], arr[maxIndex] = arr[maxIndex], arr[n-i-1]
		}
	}
}
