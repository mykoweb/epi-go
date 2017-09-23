package ch14q03a

import "testing"

func TestCountCharFreq(t *testing.T) {
	s := "bcdacebe"

	result := CountCharFreq(s)

	if result != "(a, 1), (b, 2), (c, 2), (d, 1), (e, 2)" {
		t.Errorf("CountCharFreq failed to produce correct result")
	}
}
