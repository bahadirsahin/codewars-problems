package set3

func Q70(s string, a, b int) string {
	if len(s) == 0 {
		return ""
	}

	if a > b {
		return s
	}

	if b > len(s) {
		b = len(s) - 1
	}

	ss := s[a : b+1]
	r := []rune(ss)

	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	ss = string(r)

	res := s[0:a] + ss + s[b+1:]

	return res
}
