package main

import (
	"fmt"
	"strings"
)

type person struct {
	Name     string
	Lastname string
	Age      int
	Height   float64
}

func (p *person) changeAge(age int) {
	str := strings.ToLower(p.Name)
	a := 0
	for _, val := range str {
		if val == 'a' {
			a++
		}
	}
	if a >= 2 && p.Age > 12 {
		p.Age = age
	}
}
func main() {
	user := person{"Anna", "Mathew", 25, 1.67}

	user.changeAge(49)
	fmt.Println(user)
}
