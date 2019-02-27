package main

import "fmt"

func createOrder() *Command {
	order := &Command{}
	order.fn = func(items []string) {
		for i, element := range items {
			fmt.Println(i, element)
		}
	}
	return order
}

func addItemToOrder(c *Command, item string) *Command {
	c.SetArguments(append(c.args, item))
	return c
}
