package ch06q11a

// NextPermutation takes a permutation provided as a slice, and returns the next
// permutation. If the permutation provided is the last permutation, then return
// the empty slice.
func NextPermutation(input []int) []int {
	result := make([]int, len(input))
	copy(result, input)

	i := len(result) - 2

	for result[i] > result[i+1] {
		i--
		if i < 0 {
			break
		}
	}

	if i == -1 {
		result := []int{} // result is the largest permutation
		return result
	}

	for j := len(result) - 1; j > i; j-- {
		if result[j] > result[i] {
			result[i], result[j] = result[j], result[i]
			break
		}
	}

	// sort.Ints(result[i+1 : len(result)])
	for k := i + 1; k < (len(result)-i-1)/2+i+1; k++ {
		result[k], result[len(result)-k+i] = result[len(result)-k+i], result[k]
	}

	return result
}
