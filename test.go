// Package twofer displays the name of who you are sharing with
package main

import "fmt"

// Problem -1
func shareWith(s string) string {
	if s == " " {
		s = "you"
	}
	return s
}

func Collats(n int) int {
	for i := 0; n != 1; i++ {
		if n%2 == 0 {
			return n / 2
		}
		return 3*n + 1
	}
	return n

}

func main() {
	res := Collats(12)
	fmt.Println(res)
	//name := shareWith("")
	//fmt.Printf("One for %s, one for me", name)

	//person := "zzzz"
	//
	//switch person {
	//case "Kate":
	//	fmt.Printf("One for %s, one for me", person)
	//case "Den":
	//	fmt.Printf("One for %s, one for me", person)
	//case "Anna":
	//	fmt.Printf("One for %s, one for me", person)
	//default:
	//	fmt.Println("One for you, one for me")
	//}

}
