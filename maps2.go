package main

import "fmt"

func main() {
	originalMap := map[int]string{
		1: "Alice",
		2: "Bob",
		3: "Alice",
	}

	invertMap(originalMap)

}

func invertMap(originalMap map[int]string) {
	invertedMap := make(map[string]int)
	for id, name := range originalMap {
		invertedMap[name] = id
	}

	fmt.Printf("%T => %#v \n", invertedMap, invertedMap)
}
