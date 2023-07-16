package set2

import (
	"fmt"
	"strings"
)

func Q42(sec int64) string {
	if sec == 0 {
		return ""
	}

	res := ""
	m, s := sec/60, sec%60
	h, m := m/60, m%60
	d, h := h/24, h%24
	y, d := d/365, d%365

	yStr, dStr, hStr, mStr, sStr := "", "", "", "", ""

	if y == 1 {
		yStr = fmt.Sprintf("%d year", y)
	} else if y > 0 {
		yStr = fmt.Sprintf("%d years", y)
	}

	if d == 1 {
		dStr = fmt.Sprintf("%d day", d)
	} else if d > 0 {
		dStr = fmt.Sprintf("%d days", d)
	}

	if h == 1 {
		hStr = fmt.Sprintf("%d hour", h)
	} else if h > 0 {
		hStr = fmt.Sprintf("%d hours", h)
	}

	if m == 1 {
		mStr = fmt.Sprintf("%d minute", m)
	} else if m > 0 {
		mStr = fmt.Sprintf("%d minutes", m)
	}

	if s == 1 {
		sStr = fmt.Sprintf("%d second", s)
	} else if s > 0 {
		sStr = fmt.Sprintf("%d seconds", s)
	}

	elems := []string{yStr, dStr, hStr, mStr, sStr}

	for _, e := range elems {
		if len(e) > 0 {
			res += fmt.Sprintf("%s, ", e)
		}
	}

	res = strings.TrimSpace(res)
	res = strings.TrimSuffix(res, ",")

	if strings.Count(res, ", ") > 0 {
		idx := strings.LastIndex(res, ", ")
		fmt.Println(idx)

		res = res[:idx] + " and" + res[idx+1:]
	}

	return res
}

// 0 0 0 - OK
// 0 0 1 - OK
// 0 1 0 - OK
// 0 1 1 - OK
// 1 0 0 - OK
// 1 0 1 - OK
// 1 1 0 - OK
// 1 1 1
