package set5

func Q137(n int) [][]int {
	if n == 0 {
		return [][]int{}
	}

	res := [][]int{}

	for i := 1; i <= n; i++ {
		arr := []int{}

		for j := 0; j < i; j++ {
			arr = append(arr, 1)
		}

		res = append(res, arr)
	}

	return res
}
