package set3

import (
	"fmt"
	"strings"
)

func Q87(str string, loc, num int) string {
	if len(str) == 0 {
		return ""
	}

	words := strings.Fields(str)

	if loc > len(words) {
		return ""
	}

	if num == 0 {
		num++
	}

	word := fmt.Sprintf("%s-", words[loc])
	res := strings.Repeat(word, num)
	res = strings.TrimSuffix(res, "-")

	return res
}
