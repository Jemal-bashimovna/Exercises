package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	length float64
	width  float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.width)
}

type Triangle struct {
	a, b, c float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.a * t.b
}

func (t Triangle) Perimeter() float64 {
	return t.a + t.b + t.c
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius
}
func (c Circle) Perimeter() float64 {
	return 3.14 * c.radius
}

func DescribeShape(s Shape) {
	fmt.Println("Area of shape is : ", s.Area())
	fmt.Println("Perimeter of shape is : ", s.Perimeter())
}

func main() {
	rectangle := Rectangle{5, 9}
	triangle := Triangle{4, 5, 3}
	circle := Circle{6}

	fmt.Println("Rectangle: ")
	DescribeShape(rectangle)

	fmt.Println("Triangle: ")
	DescribeShape(triangle)
	DescribeShape(circle)
}
