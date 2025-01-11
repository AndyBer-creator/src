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

	insertionSort(ar)

	fmt.Println(ar)
}

func insertionSort(ar []int) {
	// ваш код здесь
	n := len(ar)
	for i := 1; i < n; i++ {
		key := ar[i] // Текущий элемент для вставки
		j := i - 1

		// Сдвигаем элементы, которые больше ключа, на одну позицию вправо
		for j >= 0 && ar[j] > key {
			ar[j+1] = ar[j]
			j--
		}
		// Вставляем ключ на его место
		ar[j+1] = key
	}

}
