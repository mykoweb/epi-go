package ch13q02a

// CanFormPalindrome returns true if a string can be rearranged to form a
// palindrome, false otherwise.
func CanFormPalindrome(s string) bool {
	charCount := make(map[rune]int)
	arr := []rune(s)

	for _, char := range arr {
		charCount[char] += 1
	}

	oddCharCount := 0
	for _, val := range charCount {
		if (val % 2) != 0 {
			oddCharCount += 1
		}
	}

	if (len(arr) % 2) == 0 {
		return oddCharCount == 0
	} else {
		return oddCharCount <= 1
	}
}
