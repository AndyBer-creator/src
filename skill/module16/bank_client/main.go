package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type BankClient interface {
	Deposit(amount int)
	Withdrawal(amount int) error
	Balance() int
}

type Client struct {
	balance int
	mu      sync.Mutex
}

func (c *Client) Deposit(amount int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.balance += amount
}

func (c *Client) Withdrawal(amount int) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.balance < amount {
		return fmt.Errorf("недостаточно средств на счете")
	}
	c.balance -= amount
	return nil
}

func (c *Client) Balance() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.balance
}

func depositRoutine(client *Client, wg *sync.WaitGroup, output chan<- string, done <-chan struct{}) {
	defer wg.Done()
	for {
		select {
		case <-done:
			return
		default:
			amount := rand.Intn(10) + 1 // случайная сумма от 1 до 10
			client.Deposit(amount)
			output <- fmt.Sprintf("Зачислено: %d", amount)
			time.Sleep(time.Duration(rand.Intn(500)+500) * time.Millisecond) // задержка от 0.5 до 1 секунды
		}
	}
}

func withdrawalRoutine(client *Client, wg *sync.WaitGroup, output chan<- string, done <-chan struct{}) {
	defer wg.Done()
	for {
		select {
		case <-done:
			return
		default:
			amount := rand.Intn(5) + 1 // случайная сумма от 1 до 5
			err := client.Withdrawal(amount)
			if err != nil {
				output <- fmt.Sprintf("Ошибка снятия: %s", err)
			} else {
				output <- fmt.Sprintf("Снято: %d", amount)
			}
			time.Sleep(time.Duration(rand.Intn(500)+500) * time.Millisecond) // задержка от 0.5 до 1 секунды
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	client := &Client{}
	var wg sync.WaitGroup
	output := make(chan string)
	done := make(chan struct{}) // канал для завершения горутин

	// Запускаем 10 горутин для депозитов
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go depositRoutine(client, &wg, output, done)
	}

	// Запускаем 5 горутин для снятия средств
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go withdrawalRoutine(client, &wg, output, done)
	}

	// Обработка вывода
	go func() {
		for msg := range output {
			fmt.Println(msg)
		}
	}()

	// Обработка команд пользователя
	var command string
	for {
		fmt.Print("Введите команду (balance, deposit, withdrawal, exit): \n")
		fmt.Scanln(&command)

		switch command {
		case "balance":
			fmt.Printf("Текущий баланс: %d\n", client.Balance())
		case "deposit":
			var amount int
			fmt.Print("Введите сумму для зачисления: ")
			fmt.Scan(&amount)
			client.Deposit(amount)
			fmt.Printf("Зачислено: %d\n", amount)
		case "withdrawal":
			var amount int
			fmt.Print("Введите сумму для снятия: ")
			fmt.Scan(&amount)
			err := client.Withdrawal(amount)
			if err != nil {
				fmt.Printf("Ошибка снятия: %s\n", err)
			} else {
				fmt.Printf("Снято: %d\n", amount)
			}
		case "exit":
			fmt.Println("Завершение работы...")
			close(done)   // Отправляем сигнал завершения горутинам
			close(output) // Закрываем канал для завершения горутины вывода
			wg.Wait()     // Ждем завершения всех горутин
			return
		default:
			fmt.Println("Недопустимая команда. Используйте команды: balance, deposit, withdrawal, exit.")
		}
	}
}
