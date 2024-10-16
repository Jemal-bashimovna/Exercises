package main

import (
	"fmt"
)

//// problem - 2
//func (p person) areEqual(other person) bool {
//	if p.Name == other.Name && p.Age == other.Age {
//		return true
//	}
//	return false
//}

// problem - 3

type Address struct {
	City   string
	Street string
}

type Person0 struct {
	Name    string
	Age     int
	Address Address
}

func UpdateAddress(city, street string) (string, string) {
	city = "London"
	street = "MainStreet"
	return city, street

}

func main() {
	// problem - 1
	//person1 := person{"Anna", "Johnson", 25, 1.80}
	//newAge := 36
	//result1 := person.changeAge(person1, newAge)
	//fmt.Print(result1, "\n")
	//fmt.Println(person1)
	//
	//// problem - 2
	//person2 := person{"Ann", "Mathew", 25, 1.76}
	//result2 := person.areEqual(person1, person2)
	//fmt.Println(result2)

	// problem - 3
	person3 := Person0{
		Name: "Ivan",
		Age:  25,
		Address: Address{
			City:   "Moscow",
			Street: "Arbat",
		},
	}
	fmt.Printf("Real adress of %s: city: %s, street: %s\n", person3.Name, person3.Address.City, person3.Address.Street)

	newCity, newStreet := UpdateAddress(person3.Address.City, person3.Address.Street)
	fmt.Printf("Updated adress after function: new city: %s, new street: %s\n", newCity, newStreet)
}
