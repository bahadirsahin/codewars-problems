package set5

import (
	"fmt"
	"sort"
	"strings"
)

func Q150(s string) string {
	if len(s) == 0 {
		return ""
	}

	pairs := [][]string{}
	tokens := strings.Split(s, ";")

	for _, t := range tokens {
		pair := strings.Split(t, ":")

		pair[0], pair[1] = strings.ToUpper(pair[1]), strings.ToUpper(pair[0])

		pairs = append(pairs, pair)
	}

	sort.Slice(pairs, func(i, j int) bool {
		pair1, pair2 := pairs[i], pairs[j]

		c := strings.Compare(pair1[0], pair2[0])

		if c != 0 {
			return c < 0
		}

		return strings.Compare(pair1[1], pair2[1]) < 0
	})

	res := ""

	for i := 0; i < len(pairs); i++ {
		p := pairs[i]
		res += fmt.Sprintf("(%s, %s)", p[0], p[1])
	}

	return res
}
