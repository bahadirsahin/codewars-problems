package set4

import "unicode"

func Q105(str string, arr []int) string {
	if len(str) == 0 || len(arr) == 0 {
		return ""
	}

	m := map[int]bool{}

	for _, v := range arr {
		m[v] = true
	}

	res := []rune{}

	for i, r := range str {
		if !m[i] {
			res = append(res, r)

			continue
		}

		res = append(res, unicode.ToUpper(r))
	}

	return string(res)
}
