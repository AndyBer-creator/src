package calc

// Структура calculator
type calculator struct{}

// Конструктор для создания экземпляра calculator
func NewCalculator() *calculator {
	return &calculator{}
}

// Экспортируемый метод Calculate для выполнения арифметических операций
func (c *calculator) Calculate(a, b float64, sign string) float64 {
	switch sign {
	case "+":
		return c.add(a, b)
	case "-":
		return c.subtract(a, b)
	case "*":
		return c.multiply(a, b)
	case "/":
		return c.divide(a, b)
	default:
		return 0 // Или можно обработать ошибку
	}
}

// Неэкспортируемый метод сложения
func (c *calculator) add(a, b float64) float64 {
	return a + b
}

// Неэкспортируемый метод вычитания
func (c *calculator) subtract(a, b float64) float64 {
	return a - b
}

// Неэкспортируемый метод умножения
func (c *calculator) multiply(a, b float64) float64 {
	return a * b
}

// Неэкспортируемый метод деления
func (c *calculator) divide(a, b float64) float64 {
	if b == 0 {
		panic("Ошибка: Деление на ноль")
		// Или можно вернуть специальное значение, чтобы обозначить ошибку
	}
	return a / b
}
