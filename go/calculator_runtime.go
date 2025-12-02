package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	globalResult float64
	wg           sync.WaitGroup
)

func Add(a, b float64) float64 {
	globalResult = a + b
	return globalResult
}

func Sub(a, b float64) float64 {
	globalResult = a - b
	return globalResult
}
func Mul(a, b float64) float64 {
	return a * b
}

func Div(a, b float64) float64 {
	return a / b
}
func startLeakyGoroutine(ch <-chan float64) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case v := <-ch:
				_ = v
			default:
				time.Sleep(10 * time.Millisecond)
			}
		}
	}()
}
func PrintResult(ptr *float64) {
	fmt.Println("Результат:", *ptr)
}
func main() {
	ch := make(chan float64)
	startLeakyGoroutine(ch)

	for i := 0; i < 100; i++ {
		go Add(float64(i), float64(i))
		go Sub(float64(i), float64(i))
	}

	time.Sleep(1 * time.Second)

	var x *float64 = nil
	PrintResult(x)

	close(ch)
	wg.Wait()
}
