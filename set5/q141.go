package set5

import (
	"math/big"
)

func Q141(n int) *big.Int {
	a, b := big.NewInt(0), big.NewInt(1)

	for i := 0; i < n; i++ {
		a, b = b, a.Add(a, b).Add(a, b)
	}

	return a
}
