package set1

import (
	"strings"
)

func Q11(word string) string {
	if len(word) == 0 {
		return ""
	}

	word = strings.ToLower(word)

	m := map[string]int{}

	for _, r := range word {
		s := string(r)

		m[s]++
	}

	res := ""

	for _, r := range word {
		s := string(r)

		if m[s] == 1 {
			res += "("

			continue
		}

		if m[s] > 1 {
			res += ")"
		}
	}

	return res
}
