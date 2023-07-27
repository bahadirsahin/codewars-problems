package set4

import "strings"

func Q123(str string) string {
	if len(str) == 0 {
		return ""
	}

	if len(str) == 1 {
		return str
	}

	m := map[string]int{}

	for _, r := range str {
		s := string(r)
		s = strings.ToLower(s)
		m[s]++
	}

	for k, v := range m {
		if v > 1 {
			str = strings.ReplaceAll(str, k, "")
			str = strings.ReplaceAll(str, strings.ToUpper(k), "")
		}
	}

	if len(str) == 0 {
		return ""
	}

	res := str[:1]

	return res
}
