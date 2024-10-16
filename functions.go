package main

import "fmt"

// Problem 1
func rectangleArea(length, width float64) float64 {
	return length * width
}

// Problem 2
func factorial(n int) int {
	a := 1
	for i := 1; i <= n; i++ {
		a = a * i

	}
	return a
}

// Problem 3
func addIntAndFloat(x int, y float64) float64 {
	return float64(x) + y
}

func sqr(a int) int {
	return a * a
}

func main() {

	//Problem 1
	var a, b float64
	fmt.Scan(&a, &b)
	area := rectangleArea(a, b)
	fmt.Printf("Area of Rectangle %f", area)

	//Problem 2
	var n int
	fmt.Scan(&n)
	result := factorial(n)
	fmt.Printf("Factorial (%d) = %d", n, result)

	//Problem 3
	var num1 int
	var num2 float64
	fmt.Scan(&num1, &num2)
	result2 := addIntAndFloat(num1, num2)
	fmt.Println(result2)
}
