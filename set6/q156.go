package set6

import (
	"math"
	"sort"
)

func Q156(n int) int {
	if n == 0 {
		return 0
	}

	nn := n
	digits := []int{}
	perms := []int{}
	m := map[int]bool{}

	for ; nn > 0; nn /= 10 {
		digits = append(digits, nn%10)
	}

	Q156Rec(digits, 0, len(digits)-1, &perms, &m)
	sort.Ints(perms)

	if n == perms[len(perms)-1] {
		return -1
	}

	for i, v := range perms {
		if v == n {
			return perms[i+1]
		}
	}

	return 0
}

func Q156Rec(digits []int, l, r int, perms *[]int, m *map[int]bool) {
	if l == r {
		n := 0

		for i, d := range digits {
			n += (d * int(math.Pow10(len(digits)-1-i)))
		}

		if !(*m)[n] {
			*perms = append(*perms, n)
			(*m)[n] = true
		}
	}

	for i := l; i <= r; i++ {
		digits[i], digits[l] = digits[l], digits[i]
		Q156Rec(digits, l+1, r, perms, m)
		digits[i], digits[l] = digits[l], digits[i]
	}
}
