// Пусть это будет программа, в которой есть пять горутин, слушающих ввод с клавиатуры:
// при вводе сообщения с последующим нажатием клавиши Enter они реагируют на это событие,
// и каждая из них выводит обратно в консоль введённую строку:
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"unicode/utf8"
)

// Количество слушающих горутин
const amountOfGoroutines int = 5

// Текстовая команда, набираемая в консоли, для выхода из программы
const quitCommand string = "quit"

/*
Вариант реализации waitgroup  без испоьлзования стандартных библиотек
/ WaitGroup - Наша простенькая реализация объекта ожидания завершения группы горутин

	type WaitGroup struct {
	    counter int
	    // Условная переменная, через которую сам
	    // объект будет уведомлять другие горутины,
	    // что нужные им горутины завершили работу.
	    // Используется внутренний локер для потокобезопасного доступа
	    // к полю counter,
	    // так как counter - разделяемый между горутинами ресурс
	    c *sync.Cond
	}

// NewWaitGroup - создание объекта ожидания завершения группы горутин

	func NewWaitGroup() *WaitGroup {
	    return &WaitGroup{0, sync.NewCond(&sync.Mutex{})}
	}

// Add - добавление определенного количества
// горутин для ожидания
// завершения

	func (w *WaitGroup) Add(amount int) {
	    w.c.L.Lock()
	    w.counter += amount
	    w.c.L.Unlock()
	}

// Done - вызов данного метода сигнализирует о завершении работы одной из
// горутин

	func (w *WaitGroup) Done() {
	    w.c.L.Lock()
	    w.counter--
	    w.c.L.Unlock()
	    w.c.Broadcast()
	}

// Wait -  ожидание завершения работы добавленного количества горутин

	func (w *WaitGroup) Wait() {
	    defer w.c.L.Unlock()
	    w.c.L.Lock()
	    for w.counter != 0 {
	        w.c.Wait()
	    }
	}
*/
func main() {
	var message string
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
	// Для ожидания запуска слушающих горутин
	var startWg sync.WaitGroup
	// Для ожидания завершения работы всех запущенных слушающих
	// горутин
	var closeWg sync.WaitGroup
	startWg.Add(amountOfGoroutines)
	closeWg.Add(amountOfGoroutines)
	c := sync.NewCond(&sync.Mutex{})
	for i := 1; i <= amountOfGoroutines; i++ {
		go func(id int) {
			defer closeWg.Done()
			var outPutMessage string
			// Уведомляем главную горутину, что очередная слушающая
			// горутина запущена
			startWg.Done()
			for {
				c.L.Lock()
				c.Wait()
				c.L.Unlock()
				// Сравниваем введённую строку с командой выхода, не
				// учитывая регистр символов
				if strings.EqualFold(message, quitCommand) {
					return
				}
				if utf8.RuneCountInString(message) == 0 {
					outPutMessage = fmt.Sprintf("Горутина №%d обработала событие: Осуществлен ввод пустой строки", id)
				} else {
					outPutMessage = fmt.Sprintf("Горутина №%d обработала событие: Осуществлен ввод строки: \"%s\"", id, message)
				}
				fmt.Println(outPutMessage)
			}

		}(i)
	}
	// Ожидаем запуска всех слушающих горутин
	startWg.Wait()

	for {
		scanner.Scan()
		message = scanner.Text()
		fmt.Println("------------------------------------------------------------")
		c.Broadcast()
		// Сравниваем введённую строку с командой выхода, не учитывая
		// регистр символов
		if strings.EqualFold(message, quitCommand) {
			// Ждём завершение работы всех запущенных горутин
			closeWg.Wait()
			fmt.Println("Выход из программы")
			return
		}
	}
}
