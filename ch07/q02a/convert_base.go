package ch07q02a

import (
	"fmt"
	"strconv"
	"strings"
)

// The time complexity for this problem is as follows:
//
// Creating base10Num is O(n) where n is the length of strArr.
// Next, creating result takes O(logb2(base10Num)) time, and base10Num is
// upper-bounded by b1**n.
//
// Since logb2(b1**n) = n*logb2(b1) we get
// O(n + n*logb2(b1)) = O(n(1 + logb2(b1)))
func ConvertBase(s string, b1, b2 int) string {
	if len(s) < 1 {
		return ""
	}

	strArr := strings.Split(s, "")
	base10Num := 0

	for i := len(strArr) - 1; i >= 0; i-- {
		tmpNum, err := strconv.Atoi(strArr[i])
		if err != nil {
			fmt.Errorf("ConvertBase(%s, %v, %v) returned %v", s, b1, b2, err)
			return ""
		}
		base10Num += tmpNum * iPow(b1, (len(strArr)-i-1))
	}

	result := []string{}
	for base10Num != 0 {
		result = append([]string{convNum(strconv.Itoa(base10Num % b2))}, result...)
		base10Num /= b2
	}

	return strings.Join(result, "")
}

func convNum(s string) string {
	switch s {
	case "10":
		return "A"
	case "11":
		return "B"
	case "12":
		return "C"
	case "13":
		return "D"
	case "14":
		return "E"
	case "15":
		return "F"
	}
	return s
}

func iPow(a, b int) int {
	var result int = 1

	for 0 != b {
		if 0 != (b & 1) {
			result *= a

		}
		b >>= 1
		a *= a
	}

	return result
}
