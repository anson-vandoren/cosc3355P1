package main

// Stack represents the call stack of the hypothetical machine
type Stack []uint16

// Push places the given value on top of the call stack
func (s *Stack) Push(val uint16) {
	*s = append(*s, val)
}

// Pop removes and returns the top value on the call stack
func (s *Stack) Pop() (val uint16) {
	n := len(*s) - 1
	val = (*s)[n]
	*s = (*s)[:n]
	return val
}
