package set5

import (
	"sort"
	"strconv"
	"strings"
)

func Q124(str string) string {
	if len(str) == 0 {
		return ""
	}

	ws := strings.Fields(str)

	sort.Slice(ws, func(i, j int) bool {
		sumi, sumj := 0, 0

		wsi, _ := strconv.Atoi(ws[i])

		for ; wsi > 0; wsi /= 10 {
			sumi += wsi % 10
		}

		wsj, _ := strconv.Atoi(ws[j])

		for ; wsj > 0; wsj /= 10 {
			sumj += wsj % 10
		}

		if sumi == sumj {
			c := strings.Compare(ws[i], ws[j])

			return c < 0
		}

		return sumi < sumj
	})

	return strings.Join(ws, " ")
}
