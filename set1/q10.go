package set1

import "fmt"

func Q10(num uint) int {
	if num == 0 {
		return 0
	}

	cnt := 0
	bits := fmt.Sprintf("%08b", num)

	for _, b := range bits {
		if b == '1' {
			cnt++
		}
	}

	return cnt
}
