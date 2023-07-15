package set1

import (
	"sort"
	"strings"
)

func Q22(s1 string, s2 string) string {
	s := s1 + s2
	m := map[rune]int{}

	for _, r := range s {
		m[r]++
	}

	arr := []string{}

	for k, _ := range m {
		arr = append(arr, string(k))
	}

	sort.Strings(arr)

	res := strings.Join(arr, "")

	return res
}
