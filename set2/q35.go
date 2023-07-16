package set2

import "strings"

func Q35(s string) string {
	if len(s) == 0 {
		return ""
	}

	highestScore := 0
	highestWord := ""
	words := strings.Fields(s)

	for _, w := range words {
		score := GetWordScore(w)

		if score > highestScore {
			highestScore = score
			highestWord = w
		}
	}

	for _, w := range words {
		if strings.EqualFold(w, highestWord) {
			return w
		}
	}

	return ""
}

func GetWordScore(word string) int {
	m := map[rune]int{
		'a': 1,
		'b': 2,
		'c': 3,
		'd': 4,
		'e': 5,
		'f': 6,
		'g': 7,
		'h': 8,
		'i': 9,
		'j': 10,
		'k': 11,
		'l': 12,
		'm': 13,
		'n': 14,
		'o': 15,
		'p': 16,
		'q': 17,
		'r': 18,
		's': 19,
		't': 20,
		'u': 21,
		'v': 22,
		'w': 23,
		'x': 24,
		'y': 25,
		'z': 26,
	}

	if len(word) == 0 {
		return 0
	}

	score := 0

	for _, r := range word {
		score += m[r]
	}

	return score
}
