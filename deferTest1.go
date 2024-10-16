package main

import (
	"fmt"
	"log"
)

func main() {

	var num1 float64 = 15
	var num2 float64 = 4

	result := Divide(num1, num2)

	fmt.Printf("Результат %.2f\n", result)
}

func Divide(a, b float64) float64 {

	defer log.Println("Функция Divide завершена")

	defer log.Println("Функция Divide начато")
	if b == 0 {
		log.Println("Ошибка. Деления на 0")
		return 0
	}

	return a / b
}
