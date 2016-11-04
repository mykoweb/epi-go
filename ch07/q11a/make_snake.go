package ch07q11a

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
