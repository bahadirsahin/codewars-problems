package set1

import "math"

func Q26(n int) bool {
	if n < 2 {
		return false
	}

	if n == 2 || n == 3 {
		return true
	}

	sq := int(math.Sqrt(float64(n)))

	for i := 2; i <= sq; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
