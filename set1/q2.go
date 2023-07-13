package set1

import (
	"strings"
)

func Q2(s string) string {
	s = strings.Title(s)
	s = strings.ReplaceAll(s, " ", "")

	return s
}
