package set2

import (
	"strings"
	"unicode"
)

func Q56(str string) string {
	if len(str) == 0 {
		return ""
	}

	upper := 0

	for _, r := range str {
		if unicode.IsUpper(r) {
			upper++
		}
	}

	if upper > (len(str) / 2) {
		return strings.ToUpper(str)
	}

	return strings.ToLower(str)
}
