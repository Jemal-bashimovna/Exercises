package main

import "fmt"

func main() {
	var num1, num2 float64 = 15.6, 7
	result := Divide(num1, num2)
	fmt.Printf("Result: %.3f\n", result)

}

func Divide(a, b float64) float64 {
	if b == 0 {
		return 0
	}

	return a / b
}

// 	flags := log.LstdFlags | log.Lshortfile
//
//	fileInfo, _ := os.OpenFile("log_info.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
//	fileWarn, _ := os.OpenFile("log_warn.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
//	fileErr, _ := os.OpenFile("log_err.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
//
//	logInfo := log.New(fileInfo, "INFO:", flags)
//	logWarn := log.New(fileWarn, "WARNING:", flags)
//	logErr := log.New(fileErr, "ERR:", flags)
//
//	logInfo.Println("INFORMATION IS HERE")
//	logWarn.Println("WARNING INFORMATION IS HERE")
//	logErr.Println("ERROR!")
