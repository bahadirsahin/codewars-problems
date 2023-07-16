package set1

func Q30(str string) []string {
	if len(str) == 0 {
		return []string{}
	}

	if len(str)%2 == 1 {
		str += "_"
	}

	res := []string{}

	for i := 0; i < len(str)-1; i += 2 {
		s := str[i : i+2]

		res = append(res, s)
	}

	return res
}
