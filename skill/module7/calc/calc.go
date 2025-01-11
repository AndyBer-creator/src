package calc

type calculator struct{}

func NewCalculator() *calculator {
	return &calculator{}
}

func (c *calculator) Calculate(a, b float64, sign string) float64 {
	switch sign {
	case "+":
		return c.addition(a, b)
	case "-":
		return c.subtract(a, b)
	case "*":
		return c.multiply(a, b)
	case "/":
		return c.divide(a, b)
	default:
		return 0
	}
}

func (c *calculator) addition(a, b float64) float64 {
	return a + b
}

func (c *calculator) subtract(a, b float64) float64 {
	return a - b
}

func (c *calculator) multiply(a, b float64) float64 {
	return a * b
}

func (c *calculator) divide(a, b float64) float64 {
	if b == 0 {
		panic("Ошибка: Деление на ноль")

	}
	return a / b
}
