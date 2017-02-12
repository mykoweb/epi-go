package ch16q03a

// GeneratePermutations takes a slice of distinct integers and returns all the
// unique permutations.
func GeneratePermutations(i int, arr *[]int, result *[][]int) {
	if i == len(*arr)-1 {
		// Need to copy the array before appending, otherwise all the appended
		// arrays will point to the same underlying memory location.
		arrToAppend := make([]int, len(*arr))
		copy(arrToAppend, *arr)
		*result = append(*result, arrToAppend)
	}

	for j := i; j < len(*arr); j++ {
		(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
		GeneratePermutations(i+1, arr, result)

		// Need to return the ith element back to its original place, otherwise
		// we can get duplicate permutations.
		(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
	}
}
