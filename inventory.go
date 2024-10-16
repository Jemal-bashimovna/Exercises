package main

import (
	"fmt"
)

type Item struct {
	Name     string
	Category string
	Quantity int
	Price    float64
}

func addItem(inventory []Item, newItem Item) []Item {
	inventory = append(inventory, newItem)
	return inventory
}

func (i *Item) updateQuantity(newQuantity int) {
	i.Quantity = newQuantity
}
func findItemsByCategory(inventory []Item, category string) []Item {
	var newInventory []Item
	for i, val := range inventory {
		if val.Category == category {
			newInventory = append(newInventory, inventory[i])
		}
	}
	return newInventory
}

func printInventory(inventory []Item) {
	for i, val := range inventory {
		fmt.Printf("[%d] =>  Name: %s, Category: %s, Quantity %d, Price: %f\n", i, val.Name, val.Category, val.Quantity, val.Price)
	}
}

func validateTmPhoneNumber(phoneNumber string) bool {
	/*
		Example of valid phone number: +99361267156
		Example of invalid phone number: 99361267156
	*/

	if !(len(phoneNumber) == 12 && phoneNumber[0:4] == "+993") {
		return false
	}
	for i := 1; i < 12; i++ {
		if !(phoneNumber[i] >= '0' && phoneNumber[i] <= '9') {
			return false
		}
	}

	return true
}

func main() {
	var items []Item
	items = addItem(items, Item{"Fish", "seafood", 25, 29.9})
	items = addItem(items, Item{"Grape", "fruits", 25, 29.9})
	items = addItem(items, Item{"Apple", "fruits", 25, 9.99})
	fmt.Println(items)

	items[0].updateQuantity(40)
	items[0].Quantity = 48
	fmt.Println("new", items)

	items = findItemsByCategory(items, "fruits")
	fmt.Println(items)

	printInventory(items)

	validateNumber := validateTmPhoneNumber("+99361334710")
	fmt.Println(validateNumber)

}
