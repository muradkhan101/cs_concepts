package main

import (
	"fmt"
)

func main() {
	fmt.Print("Printing 4: ")
	fmt.Println(BitToString(4))
	fmt.Print("Setting bit at index 3 to 1: ")
	fmt.Println(BitToString(SetBit(4, 3)))

	fmt.Print("Printing 55: ")
	fmt.Println(BitToString(55))
	fmt.Print("Getting bit at index 3: ")
	fmt.Println(GetBit(55, 3))
	fmt.Print("Getting bit at index 4: ")
	fmt.Println(GetBit(55, 4))
}
