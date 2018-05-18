package main

import (
	"math/rand"
)

// Vector is a one dimensional collection of values
type Vector struct {
	Entries []int
}

func (v *Vector) Add(v2 *Vector) *Vector {
	if len(v.Entries) != len(v2.Entries) {
		panic("Length of Vectors doesn't match")
	}
	final := make([]int, len(v.Entries))
	for i := 0; i < len(v.Entries); i++ {
		final[i] = v.Entries[i] + v2.Entries[i]
	}
	return &Vector{final}
}

func (v *Vector) Subtract(v2 *Vector) *Vector {
	if len(v.Entries) != len(v2.Entries) {
		panic("Length of Vectors doesn't match")
	}
	final := make([]int, len(v.Entries))
	for i := 0; i < len(v.Entries); i++ {
		final[i] = v.Entries[i] - v2.Entries[i]
	}
	return &Vector{final}
}

func (v *Vector) Equals(v2 *Vector) bool {
	if len(v.Entries) != len(v2.Entries) {
		return false
	}
	for i := 0; i < len(v.Entries); i++ {
		if v.Entries[i] != v2.Entries[i] {
			return false
		}
	}
	return true
}

func (v *Vector) LessThanEqual(v2 *Vector) bool {
	if len(v.Entries) != len(v2.Entries) {
		return false
	}
	for i := 0; i < len(v.Entries); i++ {
		if v.Entries[i] > v2.Entries[i] && v.Entries[i] != v2.Entries[i] {
			return false
		}
	}
	return true
}

func (v *Vector) GreaterThanEqual(v2 *Vector) bool {
	if len(v.Entries) != len(v2.Entries) {
		return false
	}
	for i := 0; i < len(v.Entries); i++ {
		if v.Entries[i] < v2.Entries[i] && v.Entries[i] != v2.Entries[i] {
			return false
		}
	}
	return true
}

func (v *Vector) LessThan(v2 *Vector) bool {
	return v.GreaterThanEqual(v2)
}

func (v *Vector) GreaterThan(v2 *Vector) bool {
	return v.LessThanEqual(v2)
}

func RandomVector(bounds Vector) *Vector {
	newVec := Vector{make([]int, 0)}
	for i := 0; i < len(bounds.Entries); i++ {
		entry := bounds.Entries[i]
		if entry == 0 {
			newVec.Entries = append(newVec.Entries, []int{0}...)
		} else {
			newVec.Entries = append(newVec.Entries, []int{rand.Intn(entry)}...)
		}
	}
	return &newVec
}
