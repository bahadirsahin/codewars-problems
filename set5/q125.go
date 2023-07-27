package set5

import "math"

func Q125(n int) int {
	if n == 0 {
		return 0
	}

	lg := 0

	for i := 1; i < math.MaxInt; i++ {
		if int(math.Pow(float64(i), 5)) >= n {
			lg = i

			break
		}
	}

	cnt := 0

	for i := 1; i <= lg; i++ {
		cnt += n / int(math.Pow(5, float64(i)))
	}

	return cnt
}
