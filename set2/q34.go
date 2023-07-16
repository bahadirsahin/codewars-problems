package set2

func Q34(n int) []string {
	if n == 0 {
		return []string{}
	}

	res := []string{}

	for i := 1; i <= n; i++ {
		s := ""

		for j := 0; j < (n - i); j++ {
			s += " "
		}

		for j := 0; j < (2*i - 1); j++ {
			s += "*"
		}

		for j := 0; j < (n - i); j++ {
			s += " "
		}

		res = append(res, s)
	}

	return res
}
