package ch07q11a

// MakeSnake takes a string, s, and returns the snake string. A snake string is
// defined as the left-right, top-bottom representation of a string written
// sinusoidally.
//
// For example, the snake string for "Hello_World!" is "e_lHloWrdlo!".
//
// The time complexity is O(n) where n is the length of the string.
func MakeSnake(s string) (result string) {
	for i := 1; i < len(s); i += 4 {
		result = result + string(s[i])
	}
	for i := 0; i < len(s); i += 2 {
		result = result + string(s[i])
	}
	for i := 3; i < len(s); i += 4 {
		result = result + string(s[i])
	}

	return
}
