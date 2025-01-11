package main

import (
	"fmt"
)

func main() {
	var size1, size2 int

	// Запрашиваем размер первого массива
	fmt.Print("Введите размер первого массива: ")
	fmt.Scan(&size1)

	// Запрашиваем размер второго массива
	fmt.Print("Введите размер второго массива: ")
	fmt.Scan(&size2)

	// Создаем первый массив
	array1 := make([]string, size1)
	fmt.Println("Введите элементы первого массива:")
	for i := 0; i < size1; i++ {
		fmt.Print("Элемент ", i+1, ": ")
		fmt.Scan(&array1[i])
	}

	// Создаем второй массив
	array2 := make([]string, size2)
	fmt.Println("Введите элементы второго массива:")
	for i := 0; i < size2; i++ {
		fmt.Print("Элемент ", i+1, ": ")
		fmt.Scan(&array2[i])
	}

	// Находим общие элементы
	commonElements := findCommonElements(array1, array2)

	// Выводим общие элементы
	if len(commonElements) > 0 {
		fmt.Println("Общие элементы:", commonElements)
	} else {
		fmt.Println("Нет общих элементов.")
	}
}

// Функция для нахождения общих элементов в двух массивах
func findCommonElements(array1, array2 []string) []string {
	elementMap := make(map[string]bool)
	common := []string{}

	// Заполняем map элементами первого массива
	for _, element := range array1 {
		elementMap[element] = true
	}

	// Проверяем элементы второго массива на наличие в map
	for _, element := range array2 {
		if elementMap[element] {
			common = append(common, element)
		}
	}

	return common
}

// Что можно доработать:
// 1. В коде отсутствует проверка на корректность ввода размера массивов.
//Если пользователь введет отрицательное число или ноль, программа создаст
//массив с некорректным размером.
//Решение: добавить валидацию на положительный размер.
// 2. Если вводимые элементы содержат дубли в первом массиве,
//они будут включены в результат многократно.
//Решение: можно использовать map для исключения дублирующихся значений
//в итоговом списке общих элементов.
// 3. Если массивы пустые (размеры 0), имеет смысл заранее уведомлять пользователя
//о невозможности сравнения, чтобы не выполнять лишние операции.
