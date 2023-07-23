package set4

import (
	"strings"
	"unicode"
)

func Q104(str string) string {
	if len(str) == 0 {
		return ""
	}

	res := []string{}
	words := strings.Fields(str)

	for _, word := range words {
		runeArr := []rune{}

		for i, r := range word {
			if i%2 == 0 {
				runeArr = append(runeArr, unicode.ToUpper(r))
			} else {
				runeArr = append(runeArr, unicode.ToLower(r))
			}
		}

		res = append(res, string(runeArr))
	}

	return strings.Join(res, " ")
}
