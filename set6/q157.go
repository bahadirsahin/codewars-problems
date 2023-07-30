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

	arr := [][2]int{intervals[0]}

	for i, in := range intervals {
		if i == 0 {
			continue
		}

		start, end, lastEnd := in[0], in[1], arr[len(arr)-1][1]

		if start <= lastEnd {
			max := 0

			if end > lastEnd {
				max = end
			} else {
				max = lastEnd
			}

			arr[len(arr)-1][1] = max
		} else {
			arr = append(arr, in)
		}
	}

	sum := 0

	for _, in := range arr {
		sum += (in[1] - in[0])
	}

	return sum
}
