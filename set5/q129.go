package set5

import (
	"fmt"
	"math"
	"sort"
)

func Q129(n int) string {
	if n == 0 {
		return ""
	}

	primes := []int{}

	for i := 2; i < int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 && isPrime(i) {
			primes = append(primes, i)
		}
	}

	nn := n
	divsMap := map[int]int{}

	for _, p := range primes {
		for nn%p == 0 {
			divsMap[p]++
			nn /= p
		}
	}

	if nn > 1 {
		divsMap[nn]++
	}

	keys := make([]int, 0, len(divsMap))

	for k := range divsMap {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	res := ""

	for _, k := range keys {
		d := k
		c := divsMap[k]

		if c == 1 {
			res += fmt.Sprintf("(%d)", d)
		} else {
			res += fmt.Sprintf("(%d**%d)", d, c)
		}
	}

	return res
}

func isPrime(n int) bool {
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
