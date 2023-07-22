package set3

import "fmt"

func Q79(n int) []int {
	if n == 0 {
		return []int{}
	}

	res := []int{}

	for i := 1; i <= n; i++ {
		bits := fmt.Sprintf("%b", i)

		if bits[0] == '1' && bits[len(bits)-1] == '1' {
			res = append(res, i)
		}
	}

	return res
}
