package main

import "fmt"

type Operation interface {
	Compute(x, y float64) float64
}

//, Subtraction, Multiplication и Division

type Addition struct {
	x, y float64
}

type Subtraction struct {
	x, y float64
}

type Multiplication struct {
	x, y float64
}

type Division struct {
	x, y float64
}

func (a Addition) Compute(x, y float64) float64 {
	return x + y
}

func (s Subtraction) Compute(x, y float64) float64 {
	return x - y
}
func (m Multiplication) Compute(x, y float64) float64 {
	return x * y
}

func (d Division) Compute(x, y float64) float64 {
	return x / y
}

func PerformOperation(o Operation, num1, num2 float64) float64 {
	return o.Compute(num1, num2)
}

func main() {
	var addNum Operation = Addition{}
	var subNum Operation = Subtraction{}
	var multNum Operation = Multiplication{}
	var divNum Operation = Division{}

	fmt.Println(PerformOperation(addNum, 67, 6))
	fmt.Println(PerformOperation(subNum, 7, 3))
	fmt.Println(PerformOperation(multNum, 6, 6))
	fmt.Println(PerformOperation(divNum, 10, 2))

}
