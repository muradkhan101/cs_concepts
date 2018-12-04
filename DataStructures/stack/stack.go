package main

// Stack typical
type Stack struct {
	Top *Node
}

// Push to stack
func (s *Stack) Push(val int) {
	node := &Node{val, nil}
	node.Next = s.Top
	s.Top = node
}

// Pop from stack
func (s *Stack) Pop() int {
	if s.Top == nil {
		panic("Popping from empty stack")
	}
	value := s.Top.Value
	s.Top = s.Top.Next
	return value
}
