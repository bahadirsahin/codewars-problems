package set4

import "math"

func Q110(words []string, k int) string {
	if len(words) == 0 || k == 0 {
		return ""
	}

	idxMap := map[string]int{}
	valMap := map[string]int{}

	for i := 0; i < len(words)-k+1; i++ {
		word := ""

		for j := 0; j < k; j++ {
			word += words[i+j]
		}

		valMap[word] = len(word)
		idxMap[word] = i
	}

	max := math.MinInt

	for _, v := range valMap {
		if v > max {
			max = v
		}
	}

	longestWords := []string{}

	for k, v := range valMap {
		if v == max {
			longestWords = append(longestWords, k)
		}
	}

	if len(longestWords) == 1 {
		return longestWords[0]
	}

	minIdx := math.MaxInt

	for _, lw := range longestWords {
		if idxMap[lw] < minIdx {
			minIdx = idxMap[lw]
		}
	}

	for k, v := range idxMap {
		if v == minIdx {
			return k
		}
	}

	return ""
}
