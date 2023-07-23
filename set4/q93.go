package set4

import "strings"

func Q93(s string) string {
	if len(s) == 0 {
		return ""
	}

	chars := strings.Split(s, "")
	res := strings.Join(chars, " ")

	return res
}
