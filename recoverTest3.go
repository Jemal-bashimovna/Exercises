package main

import (
	"fmt"
)

func main() {
	defer func() {
		fmt.Printf("Recover. %v\n")
	}()

	fmt.Printf("Result: %.2f\n", SafeDivide(9, 0))
}

func SafeDivide(a, b float64) float64 {
	if b == 0 {
		panic("Cannot divide a number by zero")
	}
	return a / b
}
