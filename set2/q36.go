package set2

import (
	"sort"
	"strings"
)

func Q36(arr1 []string, arr2 []string) []string {
	if len(arr1) == 0 || len(arr2) == 0 {
		return []string{}
	}

	m := map[string]bool{}
	res := []string{}

	for _, w1 := range arr1 {
		for _, w2 := range arr2 {
			if strings.Contains(w2, w1) {
				if !m[w1] {
					res = append(res, w1)
					m[w1] = true
				}
			}
		}
	}

	sort.Strings(res)

	return res
}
