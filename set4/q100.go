package set4

import "math"

func Q100(a, b, c float64) (area float64) {
	s := (a + b + c) / 2
	res := s
	res *= (s - a)
	res *= (s - b)
	res *= (s - c)
	res = math.Sqrt(res)

	return res
}
