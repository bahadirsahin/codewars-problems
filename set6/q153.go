package set6

func Q153(arr [][]int) []int {
	if arr == nil || len(arr[0]) == 0 {
		return []int{}
	}

	res := []int{}
	n := len(arr)

	for i := 0; i < n/2; i++ {
		for j := i; j < n-i; j++ {
			res = append(res, arr[i][j])
		}

		for j := i + 1; j < n-i; j++ {
			res = append(res, arr[j][n-1-i])
		}

		for j := i + 1; j < n-i; j++ {
			res = append(res, arr[n-1-i][n-1-j])
		}

		for j := i + 1; j < n-i-1; j++ {
			res = append(res, arr[n-1-j][i])
		}
	}

	if n%2 == 1 {
		res = append(res, arr[n/2][n/2])
	}

	return res
}
