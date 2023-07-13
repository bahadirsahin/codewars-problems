package set1

import "strings"

func Q3(str string) string {
	if len(str) == 0 {
		return ""
	}

	resArr := []string{}
	wordMap := map[string]bool{}
	words := strings.Fields(str)

	for _, w := range words {
		if !wordMap[w] {
			resArr = append(resArr, w)
		}

		wordMap[w] = true
	}

	res := strings.Join(resArr, " ")

	return res
}
