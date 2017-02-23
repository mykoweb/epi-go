package ch09q01

// MaxStack implements a stack with a Max() function that returns the maximum
// value stored in the stack.
type MaxStack struct {
	stack    []int
	maxStack []int
}

func (s *MaxStack) Max() (int, bool) {
	if len(s.maxStack) == 0 {
		return 0, false
	}
	return s.maxStack[len(s.maxStack)-1], true
}

func (s *MaxStack) Push(val int) {
	if len(s.maxStack) == 0 || val >= s.maxStack[len(s.maxStack)-1] {
		s.maxStack = append(s.maxStack, val)
	}

	s.stack = append(s.stack, val)
}

func (s *MaxStack) Pop() (int, bool) {
	if len(s.stack) == 0 {
		return 0, false
	}

	if s.maxStack[len(s.maxStack)-1] == s.stack[len(s.stack)-1] {
		s.maxStack = s.maxStack[0 : len(s.maxStack)-1]
	}

	val := s.stack[len(s.stack)-1]
	s.stack = s.stack[0 : len(s.stack)-1]
	return val, true
}
