package ch19q02a

import "testing"

func TestPaintBoolMatrix(t *testing.T) {
	// Test the following matrix
	// false false false
	// true  false false
	// true  false true
	// where the bottom left square has x, y coords of {0, 0} and the square to its
	// right has coords of {1, 0}.
	// The initial square is {0, 0}.

	sq0_0 := &Square{X: 0, Y: 0, Val: true}
	sq1_0 := &Square{X: 1, Y: 0}
	sq2_0 := &Square{X: 2, Y: 0, Val: true}
	sq0_1 := &Square{X: 0, Y: 1, Val: true}
	sq1_1 := &Square{X: 1, Y: 1}
	sq2_1 := &Square{X: 2, Y: 1}
	sq0_2 := &Square{X: 0, Y: 2}
	sq1_2 := &Square{X: 1, Y: 2}
	sq2_2 := &Square{X: 2, Y: 2}

	(*sq0_0).Adj = [4]*Square{nil, sq1_0, sq0_1, nil}
	(*sq1_0).Adj = [4]*Square{sq0_0, sq2_0, sq1_1, nil}
	(*sq2_0).Adj = [4]*Square{sq1_0, nil, sq2_1, nil}
	(*sq0_1).Adj = [4]*Square{nil, sq1_1, sq0_2, sq0_0}
	(*sq1_1).Adj = [4]*Square{sq0_1, sq2_1, sq1_2, sq1_0}
	(*sq2_1).Adj = [4]*Square{sq1_1, nil, sq2_2, sq2_0}
	(*sq0_2).Adj = [4]*Square{nil, sq1_2, nil, sq0_1}
	(*sq1_2).Adj = [4]*Square{sq0_2, sq2_2, nil, sq1_1}
	(*sq2_2).Adj = [4]*Square{sq1_2, nil, nil, sq2_1}

	start := sq0_0

	PaintBoolMatrix(start)
	// Final matrix should look like
	// false false false
	// false false false
	// false false true
	squares := []*Square{sq0_0, sq1_0, sq2_0, sq0_1, sq1_1, sq2_1, sq0_2, sq1_2, sq2_2}
	for i, sq := range squares {
		if i == 2 {
			if sq.Val == false {
				t.Errorf("The lone false square turned out to be true")
			}
		} else {
			if sq.Val == true {
				t.Errorf("This square should be false but it came back true")
			}
		}
	}
}
