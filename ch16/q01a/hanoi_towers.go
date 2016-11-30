package ch16q01a

import epi "github.com/mykoweb/epi-go/ch09/stack"

// Move pegs from src stack to dst stack making sure that for a peg x that is on
// top of peg y within a given stack, the value for peg x is always less than
// the value for peg y.
func HanoiTowers(numRings int, src, dst, buf *epi.Stack) {
	if numRings > 0 {
		HanoiTowers(numRings-1, src, buf, dst)
		top := src.Pop()
		dst.Push(top)
		HanoiTowers(numRings-1, buf, dst, src)
	}
}
