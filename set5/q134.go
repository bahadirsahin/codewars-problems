package set5

import (
	"strconv"
	"strings"
)

func Q134(ip string) bool {
	if len(ip) == 0 {
		return false
	}

	tokens := strings.Split(ip, ".")

	if len(tokens) != 4 {
		return false
	}

	for _, t := range tokens {
		if len(t) > 1 && strings.HasPrefix(t, "0") {
			return false
		}

		td, err := strconv.Atoi(t)

		if err != nil {
			return false
		}

		if td < 0 || td > 255 {
			return false
		}
	}

	return true
}
