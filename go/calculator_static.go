package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func BadStyleFunction(a, b float64) float64 {
	return a + b
}

var bigLeak [][]byte

func Add(a, b float64) float64 {
	return a + b
}
func Sub(a, b float64) float64 {
	return a - b
}
func Mul(a, b float64) float64 {
	return a + b
}
func Div(a, b float64) float64 {
	return a / b
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

	for i := 0; i < 10; i++ {
		chunk := make([]byte, 10*1024*1024)
		bigLeak = append(bigLeak, chunk)
	}

	var result float64

	switch op {
	case "+":
		result = Add(a, b)
	case "-":
		result = Sub(a, b)
	case "*":
		result = Mul(a, b)
	case "/":
		result = Div(a, b)
	default:
		fmt.Println("Неизвестная операция")
		return
	}

	fmt.Printf("Результат: %.2f\n", result)

	_ = time.Second
}
