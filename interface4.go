package main

import (
	"fmt"
	"io/ioutil"
)

type Logger interface {
	Log(message string)
}

type ConsoleLogger struct {
	fileName string
}

func (c ConsoleLogger) Log(message string) {
	fmt.Println(c.fileName, " - ", message, " saved in Console")
}

type FileLogger struct {
	fileName string
}

func (f FileLogger) Log(message string) {
	data := []byte(message)
	err := ioutil.WriteFile(f.fileName, data, 0600)
	if err != nil {
		fmt.Println("Cannot create file")
	} else {
		fmt.Println("File successfully created")
	}
}

func TestLogger(lg Logger) {
	lg.Log("Some Message...")
}

func main() {
	var consoleLog Logger = ConsoleLogger{"text.txt"}
	var fileLog Logger = FileLogger{"text.txt"}

	TestLogger(consoleLog)

	TestLogger(fileLog)

}
