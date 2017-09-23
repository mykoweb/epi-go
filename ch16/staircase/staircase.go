package staircase

// Given a staircase with n steps, calculate how many possible ways one can run
// up the stairs if one can only run up 1 step, 2 steps, or 3 steps.
func Staircase(n int) int {
	if n < 0 {
		return 0
	} else if n == 0 {
		return 1
	}

	return Staircase(n-3) + Staircase(n-2) + Staircase(n-1)
}
