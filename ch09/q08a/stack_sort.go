package ch09q08a

import "errors"

type node struct {
	val  int
	next *node
}

type Stack struct {
	top  *node
	size int
}

func (s *Stack) Pop() (val int, err error) {
	if s.size > 0 {
		val, s.top = s.top.val, s.top.next
		s.size--
		return
	}
	return 0, errors.New("Stack underflow error")
}

func (s *Stack) Push(val int) {
	s.top = &node{val, s.top}
	s.size++
}

func (s *Stack) Peek() (int, error) {
	if s.size > 0 {
		return s.top.val, nil
	} else {
		return 0, errors.New("Stack empty")
	}
}

func (s *Stack) Len() int {
	return s.size
}

func StackSort(s *Stack) {
	if s.Len() > 0 {
		val, _ := s.Pop()
		StackSort(s)
		insert(val, s)
	}
}

func insert(val int, s *Stack) {
	topVal, _ := s.Peek()
	if s.Len() < 1 || val >= topVal {
		s.Push(val)
	} else {
		s.Pop()
		insert(val, s)
		s.Push(topVal)
	}
}
