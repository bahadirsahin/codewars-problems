package set2

import "math"

func Q32(n, p int) int {
	nn := n
	arr := []int{}

	for nn > 0 {
		arr = append(arr, nn%10)
		nn /= 10
	}

	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	sum := 0

	for i, n := range arr {
		sum += int(math.Pow(float64(n), float64(p+i)))
	}

	if sum%n != 0 {
		return -1
	}

	return sum / n
}
