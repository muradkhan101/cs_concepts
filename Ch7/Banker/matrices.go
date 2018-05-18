package main

// Matrix is two dimensional Matrix made of  Vectors
type Matrix struct {
	Entries []Vector
	Vlen    int //Length of Vectors in Entries array
}

func (m *Matrix) AddEntry(v *Vector) {
	if m.Vlen != len(v.Entries) {
		panic("Vector length mismatch")
	}
	m.Entries = append(m.Entries, []Vector{*v}...)
}

func (m *Matrix) Add(m2 *Matrix) *Matrix {
	if m.Vlen != m2.Vlen || len(m.Entries) != len(m2.Entries) {
		panic("Matrix dimensions don't match")
	}
	final := Matrix{make([]Vector, len(m.Entries)), m.Vlen}
	for i := 0; i < len(m.Entries); i++ {
		final.Entries[i] = *m.Entries[i].Add(&m2.Entries[i])
	}
	return &final
}

func (m *Matrix) Subtract(m2 *Matrix) *Matrix {
	if m.Vlen != m2.Vlen || len(m.Entries) != len(m2.Entries) {
		panic("Matrix dimensions don't match")
	}
	final := Matrix{make([]Vector, len(m.Entries)), m.Vlen}
	for i := 0; i < len(m.Entries); i++ {
		final.Entries[i] = *m.Entries[i].Subtract(&m2.Entries[i])
	}
	return &final
}

func (m *Matrix) LessThan(m2 *Matrix) bool {
	if m.Vlen != m2.Vlen || len(m.Entries) != len(m2.Entries) {
		return false
	}
	for i := 0; i < len(m.Entries); i++ {
		if m.Entries[i].GreaterThan(&m2.Entries[i]) {
			return false
		}
	}
	return true
}

func (m *Matrix) GreaterThan(m2 *Matrix) bool {
	if m.Vlen != m2.Vlen || len(m.Entries) != len(m2.Entries) {
		return false
	}
	for i := 0; i < len(m.Entries); i++ {
		if m.Entries[i].LessThan(&m2.Entries[i]) {
			return false
		}
	}
	return true
}
