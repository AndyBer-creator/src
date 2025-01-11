package main

import (
	"bufio"
	"fmt"
	"os"
	"skill/module7/calc"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите первое число: ")
	firstStr, _ := reader.ReadString('\n')
	firstNum, err := strconv.ParseFloat(firstStr[:len(firstStr)-1], 64)

	if err != nil {
		fmt.Println("Ошибка при вводе первого числа:", err)
		return
	}

	fmt.Print("Введите оператор (+, -, *, /): ")
	sign, _ := reader.ReadString('\n')
	sign = sign[:len(sign)-1]

	fmt.Print("Введите второе число: ")
	secondStr, _ := reader.ReadString('\n')
	secondNum, err := strconv.ParseFloat(secondStr[:len(secondStr)-1], 64)

	if err != nil {
		fmt.Println("Ошибка при вводе второго числа:", err)
		return
	}

	calсInstance := calc.NewCalculator()
	result := calсInstance.Calculate(firstNum, secondNum, sign)
	fmt.Printf("Результат: %.2f\n", result)
}
