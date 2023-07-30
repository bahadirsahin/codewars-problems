package set6

func Q155(str string) []string {
	perms, res := []string{}, []string{}
	m := map[string]bool{}

	Q155Rec(str, 0, len(str)-1, &perms)

	for _, p := range perms {
		if !m[p] {
			res = append(res, p)
			m[p] = true
		}
	}

	return res
}

func Q155Rec(str string, l, r int, res *[]string) {
	if l == r {
		*res = append(*res, str)

		return
	}

	for i := l; i <= r; i++ {
		rs := []rune(str)
		rs[i], rs[l] = rs[l], rs[i]

		Q155Rec(string(rs), l+1, r, res)

		rs[i], rs[l] = rs[l], rs[i]
		str = string(rs)
	}
}
