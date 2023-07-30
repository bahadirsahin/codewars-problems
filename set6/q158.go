package set6

func Q158(n int) int {
	if n == 0 {
		return 0
	}

	arr := []int{1}
	x, y := 0, 0

	for len(arr) <= n {
		a := 2*arr[x] + 1
		b := 3*arr[y] + 1

		if a > b {
			arr = append(arr, b)
			y++
		} else if a < b {
			arr = append(arr, a)
			x++
		} else {
			arr = append(arr, a)
			x++
			y++
		}
	}

	return arr[n]
}
