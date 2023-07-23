package set4

import "strings"

func Q108(s string) string {
	if len(s) == 0 {
		return ""
	}

	if !strings.Contains(s, " ") {
		r := []rune(s)

		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			r[i], r[j] = r[j], r[i]
		}

		return string(r)
	}

	lens := []int{}
	words := strings.Fields(s)

	for _, w := range words {
		lens = append(lens, len(w))
	}

	s = strings.ReplaceAll(s, " ", "")

	r := []rune(s)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	s = string(r)
	res := ""

	for _, ll := range lens {
		res += s[0:ll]
		res += " "
		s = s[ll:]
	}

	res = strings.TrimSpace(res)

	return res
}
