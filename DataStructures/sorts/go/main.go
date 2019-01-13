package main

import (
	"fmt"
)

func main() {
	arr := make([]int, 6)
	arr[0] = 1
	arr[1] = -1
	arr[2] = 4
	arr[3] = 0
	arr[4] = 3
	arr[5] = 6
	fmt.Print("Pre sorted: ")
	fmt.Println(arr)

	quickSort(&arr, 0, len(arr)-1)
	fmt.Print("Post sort: ")
	fmt.Println(arr)
}
