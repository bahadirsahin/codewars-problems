package set5

import "math"

func Q140(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	sum := 0

	for i := 2; i < len(arr); i++ {
		if IsPrime(i) {
			sum += arr[i]
		}
	}

	return sum
}

func IsPrime(n int) bool {
	if n == 2 || n == 3 {
		return true
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
