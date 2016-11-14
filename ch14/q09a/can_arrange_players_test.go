package ch14q09a

import "testing"

func TestCanArrangePlayers(t *testing.T) {
	teamA := []int{7, 5, 3, 1}
	teamB := []int{6, 4, 2, 0}

	result := CanArrangePlayers(teamA, teamB)
	if result != true {
		t.Errorf("Expected true but got false")
	}

	teamA = []int{6, 4, 2, 0}
	teamB = []int{7, 5, 3, 1}
	result = CanArrangePlayers(teamA, teamB)
	if result != true {
		t.Errorf("Expected true but got false")
	}

	teamA = []int{6, 6, 2, 0}
	result = CanArrangePlayers(teamA, teamB)
	if result != false {
		t.Errorf("Expected false but got true")
	}
}
