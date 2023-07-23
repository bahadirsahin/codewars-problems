package set4

import "strings"

func Q95(arr []string) []string {
	if len(arr) == 0 {
		return []string{}
	}

	if len(arr) == 1 {
		return arr
	}

	res := []string{}
	last := ""

	for _, s := range arr {
		if !strings.EqualFold(s, last) {
			res = append(res, s)
			last = s
		}
	}

	return res
}
