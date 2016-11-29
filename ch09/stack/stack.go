package stack

type Node struct {
	val  interface{}
	next *Node
}

type Stack struct {
	size int
	top  *Node
}

func (s *Stack) Push(val interface{}) {
	s.top = &Node{val, s.top}
	s.size++
}

func (s *Stack) Pop() interface{} {
	if s.top == nil {
		return nil
	}
	topVal := s.top.val
	s.top = s.top.next
	s.size--
	return topVal
}

func (s *Stack) Peek() (topVal interface{}, ok bool) {
	if s.top == nil {
		return
	}
	return s.top.val, true
}

func (s *Stack) Len() int {
	return s.size
}
