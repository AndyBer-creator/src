package main

import (
	"fmt"
	"strconv"
	"time"
)

// Константы для настройки буфера и интервала
const (
	bufferSize    = 5
	flushInterval = 2 * time.Second
)

// Функция фильтрации отрицательных чисел
func filterNegative(in <-chan int, out chan<- int) {
	for num := range in {
		if num >= 0 {
			out <- num
		}
	}
	close(out)
}

// Функция фильтрации чисел, не кратных 3 (исключая 0)
func filterNotMultipleOfThree(in <-chan int, out chan<- int) {
	for num := range in {
		if num%3 == 0 {
			out <- num
		}
	}
	close(out)
}

// Кольцевой буфер
type RingBuffer struct {
	buffer []int
	size   int
	head   int
	tail   int
	count  int
}

func NewRingBuffer(size int) *RingBuffer {
	return &RingBuffer{
		buffer: make([]int, size),
		size:   size,
	}
}

func (rb *RingBuffer) Push(num int) {
	rb.buffer[rb.head] = num
	rb.head = (rb.head + 1) % rb.size
	if rb.count < rb.size {
		rb.count++
	} else {
		rb.tail = (rb.tail + 1) % rb.size
	}
}

func (rb *RingBuffer) Flush() []int {
	data := []int{}
	for i := 0; i < rb.count; i++ {
		index := (rb.tail + i) % rb.size
		data = append(data, rb.buffer[index])
	}
	rb.tail = 0
	rb.head = 0
	rb.count = 0
	return data
}

// Функция для буферизации данных из канала
func bufferData(in <-chan int, out chan<- int, done chan<- struct{}) {
	rb := NewRingBuffer(bufferSize)

	ticker := time.NewTicker(flushInterval)
	defer ticker.Stop()

	go func() {
		for num := range in {
			rb.Push(num)
		}
		// После завершения ввода, отправляем оставшиеся данные
		if rb.count > 0 {
			for _, num := range rb.Flush() {
				out <- num
			}
		}
		close(out)
		done <- struct{}{}
	}()

	for {
		select {
		case <-ticker.C:
			if len(rb.Flush()) > 0 {
				for _, num := range rb.Flush() {
					out <- num
				}
			}
		}
	}
}

// Потребитель данных
func consumer(in <-chan int) {
	for num := range in {
		fmt.Printf("Получены данные: %d\n", num)
	}
}

func main() {
	// Каналы для передачи данных между стадиями
	inputChan := make(chan int)
	negativeFilterChan := make(chan int)
	multipleOfThreeChan := make(chan int)
	outputChan := make(chan int)
	done := make(chan struct{})

	// Запуск стадий
	go filterNegative(inputChan, negativeFilterChan)
	go filterNotMultipleOfThree(negativeFilterChan, multipleOfThreeChan)
	go bufferData(multipleOfThreeChan, outputChan, done)
	go consumer(outputChan)

	// Источник данных из консоли
	fmt.Println("Введите целые числа (для завершения нажмите Enter):")
	for {
		var input string
		fmt.Scanln(&input)

		if input == "" { // Завершение ввода по нажатию Enter
			break
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Ошибка: введите целое число.")
			continue
		}

		inputChan <- num
	}

	close(inputChan)
	<-done                      // Ждем завершения обработки данных
	time.Sleep(3 * time.Second) // Дождаться завершения вывода оставшихся данных
}
