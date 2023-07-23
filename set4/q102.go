package set4

func Q102(size int) [][]int {
	if size < 1 {
		return [][]int{}
	}

	res := [][]int{}

	for i := 1; i <= size; i++ {
		arr := []int{}

		for j := 1; j <= size; j++ {
			arr = append(arr, i*j)
		}

		res = append(res, arr)
	}

	return res
}
