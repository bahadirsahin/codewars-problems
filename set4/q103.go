package set4

func Q103(data string) []int {
	if len(data) == 0 {
		return []int{}
	}

	v := 0
	res := []int{}

	for _, r := range data {
		if r == 'i' {
			v++
		}

		if r == 'd' {
			v--
		}

		if r == 's' {
			v *= v
		}

		if r == 'o' {
			res = append(res, v)
		}
	}

	return res
}
