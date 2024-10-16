package main

import (
	"fmt"
)

func main() {
	numb := []int{10, 30, 25, 40, 20}
	fmt.Println("original slice: ", numb)
	sortBubble(numb, 0, 0)
	fmt.Println("Result: ", numb)
	//sortLength()
}

func sortBubble(slice []int, i, j int) {
	for i = 0; i < len(slice); i++ {
		for j = 0; j < len(slice)-1-i; j++ {
			fmt.Printf("j=%d ; j+1=%d; j<%d\n", j, j+1, len(slice)-1-i)
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
				fmt.Printf("[%d] => %d\n", slice[j], slice[j+1])
			}
			fmt.Println(slice)
		}
		fmt.Println()
	}
}

func sortLength() {
	stringSlice := []string{"aaa", "a", "bbbbb", "bb"}
	fmt.Println("Original Slice: ", stringSlice)
	fmt.Println()

	for i := 0; i < len(stringSlice); i++ {
		for j := 0; j < len(stringSlice)-1-i; j++ {
			if len(stringSlice[j]) < len(stringSlice[j+1]) {
				stringSlice[j], stringSlice[j+1] = stringSlice[j+1], stringSlice[j]
				fmt.Printf("[%s] => %s\n", stringSlice[j], stringSlice[j+1])
			}
			fmt.Println(stringSlice)
		}
		fmt.Println()
	}

	fmt.Println("Result: ", stringSlice)
}
