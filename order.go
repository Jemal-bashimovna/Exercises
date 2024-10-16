package main

import "fmt"

type Order struct {
	ID       int
	Item     string
	Quantity int
	Status   string
}

func addOrder(orders []Order, newOrder Order) []Order {
	orders = append(orders, newOrder)
	return orders
}

func updateOrderStatus(orders []Order, id int, status string) []Order {

	for i := 0; i < len(orders); i++ {
		if orders[i].ID == id {
			orders[i].Status = status
			return orders
		}
	}
	return orders
}

func getOrder(orders []Order, id int) *Order {
	for i := range orders {
		if orders[i].ID == id {
			return &orders[i]
		}
	}
	return nil
}

func getFirstActiveOrder(orders []Order) *Order {
	for i := range orders {
		if orders[i].Status == "active" {
			return &orders[i]
		}
	}
	return nil
}

func main() {
	var newOrder []Order
	fmt.Println(newOrder)

	newOrder = addOrder(newOrder, Order{1, "Phone", 3, "active"})
	newOrder = addOrder(newOrder, Order{2, "AirPods", 5, "delivered"})
	newOrder = addOrder(newOrder, Order{3, "Laptop", 6, "active"})

	for i := range newOrder {
		fmt.Printf("[%d] = > \"%#v\"\n", i, newOrder[i])
	}
	fmt.Println()

	updateStatus := updateOrderStatus(newOrder, 4, "delivered")
	if updateStatus != nil {
		fmt.Println(updateStatus)
	} else {
		fmt.Println("This order doesn't exist")
		fmt.Println(updateStatus)
	}

	getPointer := getOrder(newOrder, 2)
	if getPointer != nil {
		fmt.Println(getPointer)
	} else {
		fmt.Println("This order doesn't exist")
	}

	getActive := getFirstActiveOrder(newOrder)
	if getActive != nil {
		fmt.Println("The first active order is:", getActive)
	} else {
		fmt.Println("Active order doesn't exist")
	}
}
