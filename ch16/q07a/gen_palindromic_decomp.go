package ch16q07a

import "strings"

func GenPalindromicDecomp(input string, partialPartition *[]string, result *[][]string) {
	charArr := strings.Split(input, "")
	if len(charArr) < 1 {
		// Note: You cannot just do the following in Go
		// *result = append(*result, *partialPartition)
		// for this problem.
		//
		// As we return up the call stack, the same pointer to partialPartition is
		// being used so the value of *result will change everytime *partialPartition
		// changes. Instead, create a copy of *partialPartition and append this to
		// *result.
		tmpPartialPartition := make([]string, len(*partialPartition))
		copy(tmpPartialPartition, *partialPartition)
		*result = append(*result, tmpPartialPartition)
		return
	}

	for i := 0; i < len(charArr); i++ {
		prefix := charArr[:i+1]
		if isPalindrome(charArr[:i+1]) == true {
			*partialPartition = append(*partialPartition, strings.Join(prefix, ""))
			GenPalindromicDecomp(strings.Join(charArr[i+1:], ""), partialPartition, result)
			*partialPartition = (*partialPartition)[:len(*partialPartition)-1]
		}
	}
}

func isPalindrome(input []string) bool {
	for i := 0; i < len(input)/2; i++ {
		if input[0] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}
