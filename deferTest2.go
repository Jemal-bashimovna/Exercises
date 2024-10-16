package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file := "file.txt"
	ReadFile(file)
}

func ReadFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Printf("Ошибка при открытии файла%v\n", err)
	}

	defer file.Close()

	fileData, err := os.ReadFile(filename)
	if err != nil {
		log.Printf("Ошибка при чтении файла", err)
	}
	fmt.Println(string(fileData))
}
