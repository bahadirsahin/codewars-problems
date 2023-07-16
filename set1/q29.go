package set1

import (
	"fmt"
	"strconv"
)

func Q29(seconds int) string {
	h, m, s := 0, 0, 0
	hStr, mStr, sStr := "", "", ""

	h = seconds / 3600
	seconds -= h * 3600
	hStr = strconv.Itoa(h)

	if len(hStr) == 1 {
		hStr = fmt.Sprintf("0%s", hStr)
	}

	m = seconds / 60
	seconds -= m * 60
	mStr = strconv.Itoa(m)

	if len(mStr) == 1 {
		mStr = fmt.Sprintf("0%s", mStr)
	}

	s = seconds
	sStr = strconv.Itoa(s)

	if len(sStr) == 1 {
		sStr = fmt.Sprintf("0%s", sStr)
	}

	res := fmt.Sprintf("%s:%s:%s", hStr, mStr, sStr)

	return res
}
