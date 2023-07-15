package set1

import (
	"strings"
	"unicode"
)

func Q19(s string) string {
	if len(s) == 0 {
		return ""
	}

	s = strings.ReplaceAll(s, "-", " ")
	s = strings.ReplaceAll(s, "_", " ")
	words := strings.Fields(s)
	res := ""

	for idx := 0; idx < len(words); idx++ {
		w := words[idx]

		if idx == 0 {
			res += w

			continue
		}

		res += CapitalizeWord(w)
	}

	return res
}

func CapitalizeWord(word string) string {
	if len(word) == 0 {
		return ""
	}

	r := []rune(word)
	r[0] = unicode.ToUpper(r[0])

	return string(r)
}
