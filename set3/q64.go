package set3

import (
	"strings"
	"unicode"
)

func Q64(s string) string {
	if len(s) == 0 {
		return ""
	}

	s = strings.ReplaceAll(s, " ", "")
	res := []rune{}

	for i, r := range s {
		res = append(res, unicode.ToUpper(r))

		if i != len(s)-1 {
			res = append(res, ' ')
			res = append(res, ' ')
		}
	}

	return string(res)
}
