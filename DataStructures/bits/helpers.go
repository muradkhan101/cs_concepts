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

// BitOn sets bit to on position
func BitOn(num int, i uint) int {
	return num | (1 << i)
}

// BitOff turns off given bit
func BitOff(num int, i uint) int {
	return num & ^(1 << i)
}

// ToggleBit toggles given bit
func ToggleBit(num int, i uint) int {
	if GetBit(num, i) == 1 {
		return BitOff(num, i)
	}
	return BitOn(num, i)
}

// Substring inserts num2 as a substring to num1 at the specified range
func Substring(num1, num2 int, i, j uint) int {
	leftMask := ^0 - ((1 << j) - 1)
	rightMask := (1 << i) - 1

	combined := leftMask | rightMask

	return (num1 & combined) | (num2 << i)
}
