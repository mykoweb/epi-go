package ch14q09a

import "sort"

func CanArrangePlayers(teamA, teamB []int) bool {
	if len(teamA) != len(teamB) {
		return false
	}

	sortedTeamA := make([]int, len(teamA))
	sortedTeamB := make([]int, len(teamB))
	copy(sortedTeamA, teamA)
	copy(sortedTeamB, teamB)
	sort.Ints(sortedTeamA)
	sort.Ints(sortedTeamB)

	// Make sure the tallest team is in TeamA. If not, swap.
	if sortedTeamA[len(sortedTeamA)-1] < sortedTeamB[len(sortedTeamB)-1] {
		sortedTeamA, sortedTeamB = sortedTeamB, sortedTeamA
	}

	for i := len(sortedTeamA) - 1; i >= 0; i-- {
		if sortedTeamA[i] < sortedTeamB[i] {
			return false
		}
	}

	return true
}
