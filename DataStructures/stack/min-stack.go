package main

// MinStack keeps track of min item too
type MinStack struct {
	Top *Node
	Min *Node
}

// Push adds to stack
func (m *MinStack) Push(val int) {
	if m.Min == nil {
		m.Min = &Node{val, nil}
		m.Top = &Node{val, nil}
	} else if val < m.Min.Value {
		m.Min = &Node{val, m.Min}
		m.Top = &Node{val, m.Top}
	} else {
		m.Top = &Node{val, m.Top}
	}
}

// Pop removes item from stack
func (m *MinStack) Pop() int {
	if m.Top == nil {
		panic("Popping from empty stack")
	}
	node := m.Top
	m.Top = m.Top.Next

	if m.Min.Value == node.Value {
		m.Min = m.Min.Next
	}

	return node.Value
}

// MinValue returns min value from stack
func (m *MinStack) MinValue() int {
	if m.Min == nil {
		panic("Checking min for empty stack")
	}
	return m.Min.Value
}
