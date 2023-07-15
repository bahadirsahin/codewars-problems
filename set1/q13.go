package set1

import "strings"

func Q13(s1 string) int {
	if len(s1) == 0 {
		return 0
	}

	s1 = strings.ToLower(s1)
	m := map[rune]int{}

	for _, r := range s1 {
		m[r]++
	}

	cnt := 0

	for _, v := range m {
		if v > 1 {
			cnt++
		}
	}

	return cnt
}
