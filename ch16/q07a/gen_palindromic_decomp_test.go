package ch16q07a

import (
	"strings"
	"testing"
)

func TestGenPalindromicDecomp(t *testing.T) {
	result := new([][]string)
	var partialPartition []string
	str := "0204451881"

	GenPalindromicDecomp(str, &partialPartition, result)

	if len(*result) != 8 {
		t.Errorf("Expected number of results is 8 but got %v", len((*result)))
	}

	// First result should just be an array of all the individual numbers
	strArray := strings.Split(str, "")
	for i, char := range strArray {
		if (*result)[0][i] != char {
			t.Errorf("Expected string to be %v but got &v", char, (*result)[0][i])
		}
	}
}
