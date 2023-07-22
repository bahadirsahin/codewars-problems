package set3

import (
	"strconv"
	"strings"
)

func Q84(str string, sep string) string {
	m := map[rune]int{}

	for _, r := range str {
		m[r]++
	}

	res := []string{}

	for _, r := range str {
		res = append(res, strconv.Itoa(m[r]))
	}

	return strings.Join(res, sep)
}
