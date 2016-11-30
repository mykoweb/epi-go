package ch16q01a

import (
	"testing"

	epi "github.com/mykoweb/epi-go/ch09/stack"
)

func TestHanoiTowers(t *testing.T) {
	src := &epi.Stack{}
	dst := &epi.Stack{}
	buf := &epi.Stack{}
	for i := 4; i > 0; i-- {
		src.Push(i)
	}

	HanoiTowers(src.Len(), src, dst, buf)

	var val int
	for i := 1; i <= 4; i++ {
		val = dst.Pop().(int)
		if val != i {
			t.Errorf("Expected dst val to be %v but got %v", i, val)
		}
	}

	if src.Len() != 0 {
		t.Errorf("Expected src tower to be empty but it was not")
	}

	if buf.Len() != 0 {
		t.Errorf("Expected buf tower to be empty but it was not")
	}
}
