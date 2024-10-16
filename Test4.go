package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// ProcessInput запрашивает у пользователя ввод чисел и обрабатывает их.
func ProcessInput() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("ERROR!", r)
		}
	}()

	var sum float64
	var input string

	fmt.Println("Enter numbers (enter 'exit' to exit):")
	for {
		_, err := fmt.Scanln(&input)
		if err != nil {
			return
		}
		if input == "exit" {
			break
		}
		num, err := strconv.ParseFloat(input, 64)
		if err != nil {
			panic("Invalid value: enter only number")
		}
		sum += num
	}

	fmt.Printf("Sum of entered numbers: %.2f\n", sum)
}

func main() {
	//ProcessInput()
	//Test()
	//printNums(6)
	//printNumsStatic()
	readerLine()
}

func Test() {

	//b := '\n'
	//fmt.Println(b)

	var summ float64
	var input string
	fmt.Println("Enter numbers:")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	for {

		_, err := fmt.Scan(&input)
		if err != nil {
			log.Println(err)
		}
		if input == "/" {
			break
		}
		num, er := strconv.ParseFloat(input, 64)
		if er != nil {
			panic("Enter only numbers")
		}
		fmt.Println(num)
		summ += num
	}
	fmt.Println("Sum of numbers: ", summ)
}

func printNums(n int) {
	for i := 0; i <= n; i++ {
		defer fmt.Println(i)
	}

}

func readerLine() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	var result float64
	fmt.Println("Enter your numbers:")

	r := bufio.NewReader(os.Stdin)
	input, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Type %T value %#v\n", input, input)

	// 0. String-den artykmach zatlary ayyryan. (strings.Trim())
	// 1. [2,4,6] array almaly. (strings)
	// 3.
	// 		a) Hemmesi dogry bolsa sum-y almaly.
	// 		b) Number dal bolsa panic bermeli.

	str := strings.Trim(input, "")
	str = strings.Trim(str, "\n")

	for _, val := range str {
		fmt.Printf("Type %T => %#v\n", val, val)
		//num, er := strconv.ParseFloat(val, 64)
	}

	//fmt.Printf("type: %T, value %#v\n", numArray, numArray)
	//for _, val := range numArray {
	//	fmt.Println(val)
	//
	//	if er != nil {
	//		panic(er)
	//	}
	//	result += num
	//}
	fmt.Println("Sum of numbers: ", result)
}
