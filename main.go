package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	increment(a)
	decrement(a)


}

func increment(a int) {
	for i := 0; i <= a; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

func decrement(a int) {
	for i := a; i > 0; i-- {
		for j := i; j > 0; j-- {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}
