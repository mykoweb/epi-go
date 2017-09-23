package ch14q03a

import (
	"sort"
	"strconv"
)

type RuneSlice []rune

func (r RuneSlice) Len() int           { return len(r) }
func (r RuneSlice) Less(i, j int) bool { return r[i] < r[j] }
func (r RuneSlice) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }

// CountCharFreq counts the frequency of characters appearing within a string.
func CountCharFreq(s string) string {
	runes := RuneSlice([]rune(s))
	sort.Sort(runes)

	if len(runes) == 0 {
		return ""
	} else if len(runes) == 1 {
		return "(" + string(runes[0]) + ", 1)"
	}

	result := ""
	currCharCount := 1
	for i := 1; i < len(runes); i++ {
		if runes[i] == runes[i-1] {
			currCharCount += 1
		} else {
			result = result + "(" + string(runes[i-1]) + ", " + strconv.Itoa(currCharCount) + "), "
			currCharCount = 1
		}
	}

	return result + "(" + string(runes[len(runes)-1]) + ", " + strconv.Itoa(currCharCount) + ")"
}
