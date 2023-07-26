package set4

import "math"

func Q7(s string) rune {
	if len(s) == 0 {
		return 0
	}

	fm, lm := map[rune]int{}, map[rune]int{}

	for i, r := range s {
		if fm[r] == 0 {
			fm[r] = i + 1
			lm[r] = i + 1
		} else {
			lm[r] = i + 1
		}
	}

	max := 0
	vm := map[rune]int{}

	for _, r := range s {
		val := lm[r] - fm[r]
		vm[r] = val

		if val > max {
			max = val
		}
	}

	rr := []rune{}

	for k, v := range vm {
		if v == max {
			rr = append(rr, k)
		}
	}

	min := math.MaxInt

	for _, r := range rr {
		if int(r) < min {
			min = int(r)
		}
	}

	return rune(min)
}
