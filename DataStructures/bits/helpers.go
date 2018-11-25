package main

import (
	"strconv"
)

// BitToString converts a number to Binary Format
func BitToString(num int) string {
	str := ""
	for num != 0 && num != -1 {
		next := num & 1
		str = strconv.Itoa(next) + str
		num = num >> 1
	}
	if num == -1 {
		return "-" + str
	}
	return str
}

// GetBit get bit at index in number
func GetBit(num int, pos uint) int {
	mask := num&(1<<pos) != 0
	if mask {
		return 1
	}
	return 0
}

// SetBit sets bit at position to 1
func SetBit(num int, pos uint) int {
	mask := 1 << pos
	return num | mask
}

// Xor calculates exclusive or between two numbers
func Xor(n1, n2 int) int {
	return ^((n1 & n2) | (^n1 & ^n1))
}
