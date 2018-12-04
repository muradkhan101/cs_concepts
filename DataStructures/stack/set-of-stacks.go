package main

// ArrayStack is a stack implemented with an array
type ArrayStack struct {
	stack []int
	Size  int
}

// Init stack array
func (s *ArrayStack) Init(size int) {
	array := make([]int, size)
	s.stack = array
	s.Size = 0
}

// Push to array
func (s *ArrayStack) Push(val int) {
	s.stack = append(s.stack, val)
	s.Size++
}

// Pop item from stack
func (s *ArrayStack) Pop() int {
	if s.Size == 0 {
		panic("Popping from empty stack")
	}
	item := s.stack[s.Size-1]
	s.Size--
	return item
}
