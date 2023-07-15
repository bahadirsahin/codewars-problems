package set1

import (
	"strconv"
	"strings"
)

func Q21(sentence string) string {
	if len(sentence) == 0 {
		return ""
	}

	m := map[int]string{}
	words := strings.Fields(sentence)

	for _, w := range words {
		for i := 1; i <= 9; i++ {
			s := strconv.Itoa(i)

			if strings.Contains(w, s) {
				m[i] = w

				break
			}
		}
	}

	res := ""

	for i := 1; i <= 9; i++ {
		w := m[i]

		if len(w) > 0 {
			res += w
			res += " "
		}
	}

	res = strings.TrimSpace(res)

	return res
}
