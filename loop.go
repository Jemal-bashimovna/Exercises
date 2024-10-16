package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Problem 1
	//for i := 1; i <= 100; i++ {
	//	if i%3 == 0 {
	//		fmt.Println(i, "Fizz")
	//	} else if i%5 == 0 {
	//		fmt.Println(i, "Buzz")
	//	} else if (i%3 == 0) && (i%5 == 0) {
	//		fmt.Println(i, "FizzBuzz")
	//	}
	//}

	// Problem 2
	//var N int
	//fmt.Println("Enter your number:")
	//fmt.Scan(&N)
	//m := 0
	//n := 1
	//for j := 1; j <= N; j++ {
	//	m += j
	//	n *= j
	//	if m > 100 {
	//		fmt.Printf("%d > 100; j= %d\n", m, j)
	//		break
	//	}
	//}
	//fmt.Printf("The sum of numbers:%d, product of numbers: %d", m, n)

	//Problem 3
	//for i := 1; i <= 10; i++ {
	//	for j := 1; j <= 10; j++ {
	//		res := i * j
	//		fmt.Printf("%d x %d = %d \t", j, i, res)
	//	}
	//	fmt.Println()
	//}

	// Problem 4
	//
	//var h int
	//fmt.Println("Enter height ")
	//fmt.Scan(&h)
	//
	//for i := 1; i <= h; i++ {
	//	for j := 1; j <= h; j++ {
	//		if j >= h-(i-1) {
	//			fmt.Print("*")
	//		} else {
	//			fmt.Print(" ")
	//		}
	//	}
	//	fmt.Println()
	//}

	// Problem 5
	//var N int
	//fmt.Println("Enter value of N: ")
	//fmt.Scan(&N)
	//if N < 8 {
	//	fmt.Println("Error")
	//	return
	//}
	//for i := 1; i <= N; i++ {
	//	for j := 1; j <= N; j++ {
	//		if i%2 == 0 {
	//			if j%2 == 0 {
	//				fmt.Print("W ")
	//			} else {
	//				fmt.Print("B ")
	//			}
	//		} else {
	//			if j%2 == 0 {
	//				fmt.Print("B ")
	//			} else {
	//				fmt.Print("W ")
	//			}
	//		}
	//	}
	//	fmt.Println()
	//}

	result := problem(15)
	fmt.Println(result)

}

func problem(a int) string {
	s := ""
	if a%3 == 0 {
		s = s + "Fizz"
	}
	if a%5 == 0 {
		s = s + "Buzz"
	}
	if s == "" {
		s = strconv.Itoa(a)
	}
	return s
}
