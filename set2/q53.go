package set2

import "unicode"

func Q53(s string) []int {
	if len(s) == 0 {
		return []int{}
	}

	upper, lower, num, spec := 0, 0, 0, 0

	res := []int{}

	for _, r := range s {
		if unicode.IsUpper(r) {
			upper++
		} else if unicode.IsLower(r) {
			lower++
		} else if unicode.IsNumber(r) {
			num++
		} else {
			spec++
		}
	}

	res = append(res, upper, lower, num, spec)

	return res
}
