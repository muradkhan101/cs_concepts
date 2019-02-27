package main

import "fmt"

func main() {
	order1 := createOrder()
	addItemToOrder(order1, "French Fried")
	addItemToOrder(order1, "Burger")

	order2 := createOrder()
	addItemToOrder(order2, "Hummus")
	addItemToOrder(order2, "Sweet Potato Fries")
	addItemToOrder(order2, "Falafel")

	fmt.Println("Placing order for items:")
	order1.Execute()
	fmt.Println()
	fmt.Println("Placing order for items:")
	order2.Execute()
}
