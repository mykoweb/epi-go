package ch07q04a

// ReplaceAndRemove takes a string and replaces all occurrences of 'a' with 'dd'
// and removes all occurrences of 'b'.
func ReplaceAndRemove(s string) string {
	runes := []rune(s)
	strLength := len(runes)

	for i := strLength - 1; i >= 0; i-- {
		if runes[i] == 'a' {
			runes = append(runes[0:i], append([]rune("dd"), runes[i+1:]...)...)
		} else if runes[i] == 'b' {
			runes = append(runes[0:i], runes[i+1:]...)
		}
	}

	return string(runes)
}
