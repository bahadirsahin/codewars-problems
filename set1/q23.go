package set1

import "math"

func Q23(num int64) int64 {
	sq := int(math.Sqrt(float64(num)))

	if (sq * sq) != int(num) {
		return -1
	}

	sq++

	return int64(sq * sq)
}
