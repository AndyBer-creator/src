package main

import (
	"fmt"
	"math/rand"
	"time"
)

// checkSliceIsSorted проверяет, отсортирован ли слайс a по возрастанию.
func checkSliceIsSorted(a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			return false
		}
	}
	return true
}

// bubbleSort сортирует слайс a с использованием алгоритма пузырьковой сортировки.
func bubbleSort(a []int) {
	n := len(a)
	if n <= 1 {
		return // Если слайс пустой или содержит один элемент, ничего не делаем
	}

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if a[j] > a[j+1] {
				// Меняем местами элементы
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}

func main() {
	// Установка генератора случайных чисел
	rand.Seed(time.Now().Unix())

	// Набор тестовых значений
	testSlices := [][]int{
		{0, 1, 2, 3, 4, 5},
		{9, 7, 4, 1, 3, 5},
		{0},
		{},
		{1, 1},
		{3, 2, 1},
		{5, 15, 2, 13, 7, 16, 10, 2},
		{1, 9, 7, 4, 6, 2, 1, 13, 22, -3, 12, 76},
	}

	// Генерация дополнительных тестовых значений
	for i := 0; i < 5; i++ {
		randomSlice := make([]int, rand.Intn(10)+1) // Слайс длиной от 0 до 10
		for j := range randomSlice {
			randomSlice[j] = rand.Intn(100) // Генерация случайных чисел от 0 до 99
		}
		testSlices = append(testSlices, randomSlice)
	}

	// Тестирование сортировки
	for _, slice := range testSlices {
		fmt.Printf("Исходный слайс: %v\n", slice)
		bubbleSort(slice)
		fmt.Printf("Отсортированный слайс: %v\n", slice)
		fmt.Printf("Слайс отсортирован: %v\n\n", checkSliceIsSorted(slice))
	}
}
