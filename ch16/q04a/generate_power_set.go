package ch16q04a

// GeneratePowerSet generates the power set for a given slice with distinct
// elements. The time complexity is O(n * 2^n).
func GeneratePowerSet(arr []int, powerSet *[][]int) {
	if len(arr) == 0 {
		*powerSet = append(*powerSet, []int{})
		return
	}

	GeneratePowerSet(arr[0:len(arr)-1], powerSet)

	// Golang's builtin copy does not do deep copies of multidimensional slices.
	for _, elem := range *powerSet {
		newElem := []int{}
		newElem = append(newElem, elem...) // for some reason, copy did not work
		newElem = append(newElem, arr[len(arr)-1])
		*powerSet = append(*powerSet, newElem)
	}
}
