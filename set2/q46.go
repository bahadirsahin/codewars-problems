package set2

import "strings"

func Q46(a int, b int, operator string) int {
	if strings.EqualFold(operator, "add") {
		return a + b
	}

	if strings.EqualFold(operator, "subtract") {
		return a - b
	}

	if strings.EqualFold(operator, "multiply") {
		return a * b
	}

	if strings.EqualFold(operator, "divide") {
		return a / b
	}

	return 0
}
