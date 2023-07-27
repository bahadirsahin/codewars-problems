package set4

import "math"

func Q127(arr []int) int {
	if len(arr) == 0 {
		return 1
	}

	if len(arr) == 1 {
		return arr[0] % 10
	}

	if len(arr) == 2 {
		b, e := arr[0], arr[1]

		if e == 1 {
			return b % 10
		}

		if b == 0 && e == 0 {
			return 1
		}

		if b == 0 && e > 0 {
			return 0
		}

		if b > 0 && e == 0 {
			return 1
		}

		b %= 10

		if e%4 == 0 {
			e = 4
		} else {
			e %= 4
		}

		return int(math.Pow(float64(b), float64(e))) % 10
	}

	b, e := arr[len(arr)-2], arr[len(arr)-1]

	if b == 0 && e == 0 {
		arr = arr[:len(arr)-2]
		arr = append(arr, 1)

		return Q127(arr)
	}

	if b == 0 && e > 0 {
		arr = arr[:len(arr)-2]
		arr = append(arr, 0)

		return Q127(arr)
	}

	if b > 0 && e == 0 {
		arr = arr[:len(arr)-2]
		arr = append(arr, 1)

		return Q127(arr)
	}

	r := math.Pow(float64(b), float64(e))

	arr = arr[:len(arr)-2]
	arr = append(arr, int(r))

	return Q127(arr)
}
