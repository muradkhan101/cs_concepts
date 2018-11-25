package main

// GetLeftIndex gets left child index for heap node
func GetLeftIndex(n int) int {
	return 2*n + 1
}

// GetRightIndex gets right child index for heap node
func GetRightIndex(n int) int {
	return 2*n + 2
}

// GetParentIndex gets parent index for heap node
func GetParentIndex(n int) int {
	return (n - 1) / 2
}

// Swap two elements in slice
func Swap(array []int, i, j int) {
	temp := array[i]
	array[i] = array[j]
	array[j] = temp
}
