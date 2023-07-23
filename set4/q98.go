package set4

import (
	"strconv"
	"strings"
)

func Q98(n int) string {
	if n == 1 {
		return "1"
	}

	arr := []string{
		strconv.Itoa(n),
	}

	for n > 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n *= 3
			n += 1
		}

		arr = append(arr, strconv.Itoa(n))
	}

	return strings.Join(arr, "->")
}
