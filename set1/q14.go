package set1

import (
	"fmt"
	"strings"
	"unicode"
)

func Q14(str string) string {
	if len(str) == 0 {
		return ""
	}

	res := ""
	words := strings.Fields(str)

	for _, w := range words {
		r := []rune(w)
		r[0] = unicode.ToUpper(r[0])
		w = string(r)

		res += fmt.Sprintf("%s ", w)
	}

	res = strings.TrimSpace(res)

	return res
}
