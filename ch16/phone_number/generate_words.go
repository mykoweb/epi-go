package phone_number

// Given a string of numbers representing a phone number, generate all the
// words that can be generated from that phone number. The words do not have to
// be dictionary words. Assume characters are only numbers and do not include
// 0 or 1.
func GenerateWords(s string) (result []string) {
	var chars string

	if len(s) == 1 {
		chars = getChars(s)
		for i := 0; i < len(chars); i++ {
			result = append(result, string(chars[i]))
		}
		return
	}

	tmpResult := GenerateWords(s[1:len(s)])
	chars = getChars(string(s[0]))
	for _, ch := range chars {
		for i := 0; i < len(tmpResult); i++ {
			result = append(result, string(ch)+tmpResult[i])
		}
	}

	return
}

func getChars(s string) string {
	switch s {
	case "2":
		return "abc"
	case "3":
		return "def"
	case "4":
		return "ghi"
	case "5":
		return "jkl"
	case "6":
		return "mno"
	case "7":
		return "pqrs"
	case "8":
		return "tuv"
	case "9":
		return "wxyz"
	}

	return ""
}
