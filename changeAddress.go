package main

import "fmt"

type Address struct {
	City   string
	Street string
}

type User struct {
	Name    string
	Age     int
	Address Address
}

func (u *User) UpdateAddress() {
	u.Address = Address{"London", "MainStreet"}

}

func main() {
	user := User{
		Name: "Ivan",
		Age:  25,
		Address: Address{
			City:   "Moscow",
			Street: "Arbat",
		},
	}
	fmt.Println(user.Address)
	user.UpdateAddress()
	fmt.Println(user)

}
