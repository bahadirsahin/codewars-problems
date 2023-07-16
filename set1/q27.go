package set1

import (
	"strings"
)

func Q27(str string) string {
	words := strings.Fields(str)
	revs := []string{}

	for _, w := range words {
		r := RevString(string(w))
		revs = append(revs, r)
	}

	res := strings.Join(revs, " ")

	if len(res) == len(str) {
		return res
	}

	for idx := 0; idx < len(str); idx++ {
		if str[idx] == ' ' && res[idx] != ' ' {
			res = res[:idx] + " " + res[idx:]
		}
	}

	return res
}
