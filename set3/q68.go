package set3

import "strings"

func Q68(st string) int {
	if len(st) == 0 {
		return 0
	}

	score := 0

	for _, r := range st {
		s := strings.ToUpper(string(r))
		score += DictScores[s]
	}

	return score
}
