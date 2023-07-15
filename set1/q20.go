package set1

func Q20(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	zeroCnt := 0

	for _, v := range arr {
		if v == 0 {
			zeroCnt++
		}
	}

	res := []int{}

	for _, v := range arr {
		if v != 0 {
			res = append(res, v)
		}
	}

	for i := 0; i < zeroCnt; i++ {
		res = append(res, 0)
	}

	return res
}
