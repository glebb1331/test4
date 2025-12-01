package main

import (
	"bufio"
	"fmt"
	"os"
)

func Add(a, b float64) float64 {
	return a + b
}
func Sub(a, b float64) float64 {
	return a - b
}
func Mul(a, b float64) float64 {
	return a * b
}
func Div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}
func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите первое число: ")
	var a float64
	fmt.Fscan(reader, &a)

	fmt.Print("Введите операцию (+ - * /): ")
	var op string
	fmt.Fscan(reader, &op)

	fmt.Print("Введите второе число: ")
	var b float64
	fmt.Fscan(reader, &b)

	var result float64
	var err error

	switch op {
	case "+":
		result = Add(a, b)
	case "-":
		result = Sub(a, b)
	case "*":
		result = Mul(a, b)
	case "/":
		result, err = Div(a, b)
	default:
		fmt.Println("Неизвестная операция")
		return
	}
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Printf("Результат: %.2f\n", result)
}
