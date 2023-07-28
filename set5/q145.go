package set5

import (
	"math"
	"strings"
)

func Q145(str string) int {
	if len(str) == 0 {
		return 0
	}

	cm := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8, "i": 9, "j": 10, "k": 11, "l": 12, "m": 13, "n": 14, "o": 15, "p": 16, "q": 17, "r": 18, "s": 19, "t": 20, "u": 21, "v": 22, "w": 23, "x": 24, "y": 25, "z": 26}
	vowels := []string{"a", "e", "i", "o", "u"}

	for _, v := range vowels {
		str = strings.ReplaceAll(str, v, " ")
	}

	str = strings.TrimSpace(str)

	if len(str) == 0 {
		return 0
	}

	tokens := strings.Fields(str)
	max := math.MinInt

	for _, t := range tokens {
		s := 0

		for _, r := range t {
			s += cm[string(r)]
		}

		if s > max {
			max = s
		}
	}

	return max
}
