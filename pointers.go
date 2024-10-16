package main

import "fmt"

func swap(a, b *int) {
	c := *a
	*a = *b
	*b = c
}

// N-2
type Person struct {
	Name *string
	Age  int
}

func updatePerson(p Person, name *string, age int) {
	p.Name = name
	p.Age = age
}

// N-3
func update(x *int, y *int) {
	//*x = *x + 5
	//*y = *y - 10
	fmt.Printf("type: %T, value %#v\n", x, y)
}

func reset(x int, y int) (*int, *int) {
	fmt.Println(x, y)
	x = 10
	y = 30
	fmt.Println(x, y)
	a := &x
	b := &y
	fmt.Println(a, b)
	return &x, &y
}
func main() {
	// N-1
	//var x, y int = 5, 10
	//swap(&x, &y)
	//fmt.Println(x, y)

	//N-2
	//var name string = "Alice"
	//
	//person := Person{Name: &name, Age: 25}
	//
	//fmt.Printf("Type: %T, Value: %#v\n", person, person)
	//
	//var a string = "Bob"
	//
	//updatePerson(person, &a, 30)
	//
	//fmt.Printf("Name = %s, Age = %d\n", *person.Name, person.Age)

	// N-3

	x := 20
	y := 60

	fmt.Println("maxConnection: ", x)
	fmt.Println("timeOut: ", y)

	update(&x, &y)
	fmt.Println("maxConnection: ", x)
	fmt.Println("timeOut: ", y)
	fmt.Printf("type: %T, value %#v\n", &x, &y)

	reset(x, y)
	fmt.Println("maxConnection: ", x)
	fmt.Println("timeOut: ", y)

}
