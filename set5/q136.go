package set5

import (
	"fmt"
	"math"
	"strconv"
)

const PI = 3.141592653589793

func Q136(r float64, n int) float64 {
	// a := n * r^2 * 0.5 * sin(2 * pi / n)
	a := float64(1)
	a *= float64(n)
	a *= r
	a *= r
	a *= 0.5
	a *= math.Sin(2 * PI / float64(n))
	res := fmt.Sprintf("%.3f", a)
	resf, _ := strconv.ParseFloat(res, 64)

	return resf
}
