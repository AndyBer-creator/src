package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"skill/module6/calc"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Чтение первого числа
	fmt.Print("Введите первое число: ")
	firstInput, _ := reader.ReadString('\n')
	firstNum, _ := strconv.ParseFloat(firstInput[:len(firstInput)-1], 64)

	// Чтение оператора
	fmt.Print("Введите оператор (+, -, *, /): ")
	operator, _ := reader.ReadString('\n')
	operator = operator[:len(operator)-1] // Убираем символ новой строки

	// Чтение второго числа
	fmt.Print("Введите второе число: ")
	secondInput, _ := reader.ReadString('\n')
	secondNum, _ := strconv.ParseFloat(secondInput[:len(secondInput)-1], 64)

	// Создание экземпляра calculator
	calcInstance := calc.NewCalculator()

	// Вычисление результата
	result := calcInstance.Calculate(firstNum, secondNum, operator)

	// Вывод результата
	fmt.Printf("Результат: %.2f\n", result)
}
