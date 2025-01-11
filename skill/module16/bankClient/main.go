package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// BankClient представляет собой клиента банка
type BankClient struct {
	mu      sync.Mutex    // Мьютекс для синхронизации доступа к балансу
	balance int           // Баланс счета
	doneCh  chan struct{} // Канал для остановки горутин
}

// NewBankClient создает новый экземпляр BankClient
func NewBankClient() *BankClient {
	return &BankClient{
		balance: 0,
		doneCh:  make(chan struct{}),
	}
}

// Deposit зачисляет указанную сумму на счет
func (c *BankClient) Deposit(amount int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.balance += amount
	fmt.Printf("Deposit: %d\n", amount)
}

// Withdrawal снимает указанную сумму со счета
func (c *BankClient) Withdrawal(amount int) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.balance >= amount {
		c.balance -= amount
		fmt.Printf("Withdrawal: %d\n", amount)
		return nil
	}
	return fmt.Errorf("insufficient funds")
}

// Balance возвращает текущий баланс счета
func (c *BankClient) Balance() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.balance
}
func main() {
	rand.Seed(time.Now().UnixNano())
	client := NewBankClient()

	// Запускаем 10 горутин для зачисления средств
	for i := 0; i < 10; i++ {
		go func() {
			ticker := time.NewTicker(time.Duration(rand.Intn(500)+500) * time.Millisecond)
			defer ticker.Stop()
			for {
				select {
				case <-client.doneCh:
					return
				case <-ticker.C:
					depositAmount := rand.Intn(9) + 1
					client.Deposit(depositAmount)
				}
			}
		}()
	}

	// Запускаем 5 горутин для снятия средств
	for i := 0; i < 5; i++ {
		go func() {
			ticker := time.NewTicker(time.Duration(rand.Intn(500)+500) * time.Millisecond)
			defer ticker.Stop()
			for {
				select {
				case <-client.doneCh:
					return
				case <-ticker.C:
					withdrawalAmount := rand.Intn(4) + 1
					err := client.Withdrawal(withdrawalAmount)
					if err != nil {
						fmt.Println(err)
					}
				}
			}
		}()
	}

	// Основной цикл обработки команд пользователя
	var input string
	for {
		fmt.Scanln(&input)
		switch input {
		case "balance":
			fmt.Printf("Balance: %d\n", client.Balance())
		case "deposit":
			var amount int
			fmt.Print("Enter deposit amount: ")
			fmt.Scanf("%d", &amount)
			client.Deposit(amount)
		case "withdrawal":
			var amount int
			fmt.Print("Enter withdrawal amount: ")
			fmt.Scanf("%d", &amount)
			err := client.Withdrawal(amount)
			if err != nil {
				fmt.Println(err)
			}
		case "exit":
			close(client.doneCh)
			time.Sleep(100 * time.Millisecond) // Ждем завершения горутин
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Unsupported command. You can use commands: balance, deposit, withdrawal, exit.")
		}
	}
}
