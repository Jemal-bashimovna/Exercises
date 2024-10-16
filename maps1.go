package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("write your text: ")
	input := Scan1()
	fmt.Println("write your key")
	key := Scan1()
	result := WordCount(input, key)
	fmt.Printf("%T => %#v\n", result, result)
}

func WordCount(input, key string) map[string]int {
	wordCount := make(map[string]int)
	str := strings.Split(input, " ")
	for _, val := range str {
		if val == key {
			wordCount[val]++
		}
	}
	return wordCount
}

func Scan1() string {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	if err := input.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
	}
	return input.Text()

}

func Scan2() {
	in := bufio.NewReader(os.Stdin)
	str, err := in.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка ввода: ", err)
	}
	result := strings.Split(str, "")
	fmt.Printf("%#v\n", result)

}
