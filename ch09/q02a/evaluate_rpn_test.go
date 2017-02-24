package ch09q02a

import "testing"

func TestEvaluateRpn(t *testing.T) {
	expr := []string{"3", "4", "+", "2", "*", "1", "+", "15", "/"}

	result := EvaluateRpn(expr)

	if result != 1 {
		t.Errorf("Expected result to be 1.0 but got %v", result)
	}
}
