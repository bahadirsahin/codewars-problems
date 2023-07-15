package set1

import "strings"

func Q9(str string) string {
	if len(str) == 0 {
		return ""
	}

	words := strings.Fields(str)

	for idx := 0; idx < len(words); idx++ {
		w := words[idx]

		if len(w) >= 5 {
			words[idx] = RevString(w)
		}
	}

	res := strings.Join(words, " ")

	return res
}

func RevString(str string) string {
	if len(str) == 0 {
		return ""
	}

	res := ""

	for i := len(str) - 1; i >= 0; i-- {
		res += string(str[i])
	}

	return res
}
