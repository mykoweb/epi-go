package ch09q02a

import "strconv"

// EvaluateRpn evaluates a string slice represented in Reverse Polish Notation
// and returns the result.
func EvaluateRpn(s []string) float64 {
	stack := []float64{}
	var x, y float64

	for _, elem := range s {
		if elem == "+" || elem == "-" || elem == "*" || elem == "/" {
			y, stack = stack[len(stack)-1], stack[0:len(stack)-1]
			x, stack = stack[len(stack)-1], stack[0:len(stack)-1]
			if elem == "+" {
				stack = append(stack, x+y)
			} else if elem == "-" {
				stack = append(stack, x-y)
			} else if elem == "*" {
				stack = append(stack, x*y)
			} else {
				stack = append(stack, x/y)
			}
		} else {
			val, _ := strconv.ParseFloat(elem, 64)
			stack = append(stack, val)
		}
	}

	return stack[len(stack)-1]
}
