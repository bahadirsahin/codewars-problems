package set3

func Q66(s string) int {
	if len(s) == 0 {
		return 0
	}

	res := 0

	for _, r := range s {
		i := int(r)

		for i > 0 {
			r := i % 10

			if r == 7 {
				res += 6
			}

			i /= 10
		}
	}

	return res
}
