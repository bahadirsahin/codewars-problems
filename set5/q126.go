package set4

import (
	"math"
	"strconv"
	"strings"
)

func Q126(n1, n2 string) int {
	if len(n1) == 0 || len(n2) == 0 {
		return 0
	}

	if strings.EqualFold(n1, "0") || strings.EqualFold(n2, "0") {
		return 1
	}

	bStr, eStr := n1[len(n1)-1:], ""

	if len(n2) < 2 {
		eStr = n2
	} else {
		eStr = n2[len(n2)-2:]
		eStr = strings.TrimPrefix(eStr, "0")
	}

	b, _ := strconv.Atoi(bStr)

	if b == 0 {
		return 0
	}

	e, _ := strconv.Atoi(eStr)

	if e%4 == 0 {
		return int(math.Pow(float64(b), 4)) % 10
	}

	return int(math.Pow(float64(b), float64(e%4))) % 10
}
