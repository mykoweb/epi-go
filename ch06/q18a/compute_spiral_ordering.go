package ch06q18a

// ComputeSpiralOrdering returns the spiral ordering for an nxn matrix of ints.
func ComputeSpiralOrdering(matrix [][]int) []int {
	n := len(matrix[0])
	result := []int{}

	for offset := 0; offset <= n/2; offset++ {
		// We have a matrix where n is odd and we are at the center of the matrix
		if offset == n-offset-1 {
			result = append(result, matrix[offset][offset])
		}

		// left to right
		for i := offset; i < n-offset-1; i++ {
			result = append(result, matrix[offset][i])
		}

		// top to bottom
		for i := offset; i < n-offset-1; i++ {
			result = append(result, matrix[i][n-offset-1])
		}

		// right to left
		for i := n - offset - 1; i > offset; i-- {
			result = append(result, matrix[n-offset-1][i])
		}

		// bottom to top
		for i := n - offset - 1; i > offset; i-- {
			result = append(result, matrix[i][offset])
		}
	}

	return result
}
