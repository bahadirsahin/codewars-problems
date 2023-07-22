package set3

import (
	"fmt"
	"strings"
)

func Q85(n float64) bool {
	str := fmt.Sprintf("%v", n)

	return strings.EqualFold(str, "-0")
}
