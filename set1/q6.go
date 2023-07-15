package set1

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Q6(in string) string {
	if len(in) == 0 {
		return ""
	}

	min, max := math.MaxInt32, math.MinInt32
	numStrs := strings.Fields(in)

	for _, numStr := range numStrs {
		num, _ := strconv.Atoi(numStr)

		if num < min {
			min = num
		}

		if num > max {
			max = num
		}
	}

	return fmt.Sprintf("%d %d", max, min)
}
