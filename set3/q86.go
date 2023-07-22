package set3

import (
	"sort"
	"strings"
)

func Q86(s string) bool {
	if len(s) < 2 {
		return true
	}

	chars := strings.Split(s, "")
	sort.Strings(chars)
	s = strings.Join(chars, "")

	for i := 0; i < len(s)-1; i++ {
		if int(s[i+1])-int(s[i]) != 1 {
			return false
		}
	}

	return true
}
