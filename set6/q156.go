package set6

import (
	"sort"
	"strconv"
	"strings"
)

func Q156(n int) int {
	nn := n

	if nn == 0 {
		return 0
	}

	if nn/10 == 0 {
		return -1
	}

	m := map[int]bool{}

	for ; nn > 0; nn /= 10 {
		d := nn % 10

		if !m[d] {
			m[d] = true
		}
	}

	if len(m) == 1 {
		return -1
	}

	str := strconv.Itoa(n)

	allDesc := true

	for i := 0; i < len(str)-1; i++ {
		ld, _ := strconv.Atoi(string(str[i]))
		rd, _ := strconv.Atoi(string(str[i+1]))

		if rd > ld {
			allDesc = false
		}
	}

	if allDesc {
		return -1
	}

	res, pivot, pivotIdx, pivotSub := 0, 0, 0, 0

	for i := len(str) - 1; i > 0; i-- {
		l, r := 0, 0
		l, _ = strconv.Atoi(string(str[i-1]))
		r, _ = strconv.Atoi(string(str[i]))

		if l < r {
			pivot = l
			pivotIdx = i - 1

			break
		}
	}

	rightPart := str[pivotIdx+1:]
	rightPartArr := strings.Split(rightPart, "")
	rightPartIntArr := []int{}

	for _, v := range rightPartArr {
		vd, _ := strconv.Atoi(v)

		rightPartIntArr = append(rightPartIntArr, vd)
	}

	sort.Ints(rightPartIntArr)

	for _, v := range rightPartIntArr {
		if v > pivot {
			pivotSub = v

			break
		}
	}

	for i, v := range rightPartIntArr {
		if v == pivotSub {
			rightPartIntArr[i] = pivot

			break
		}
	}

	resStr := str[:pivotIdx]
	resStr += strconv.Itoa(pivotSub)

	for _, v := range rightPartIntArr {
		resStr += strconv.Itoa(v)
	}

	res, _ = strconv.Atoi(resStr)

	return res
}
