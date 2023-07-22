package set3

import (
	"math/big"
)

func Q71(n int) string {
	if n == 0 {
		return "1"
	}

	bi := new(big.Int).Binomial(2*int64(n), int64(n))

	return bi.String()
}
