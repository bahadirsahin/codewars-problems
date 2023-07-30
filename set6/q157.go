package set6

import (
	"sort"
)

func Q157(intervals [][2]int) int {
	if len(intervals) == 0 {
		return 0
	}

	if len(intervals) == 1 {
		return (intervals[0][1] - intervals[0][0])
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	sum := 0
	x := intervals[0][0]

	for _, i := range intervals {
		if i[1] >= x {
			max := 0

			if x > i[0] {
				max = x
			} else {
				max = i[0]
			}

			sum += i[1] - max
			x = i[1]
		}
	}

	return sum
}
