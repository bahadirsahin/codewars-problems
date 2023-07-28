package set5

import "strings"

func Q148(roman string) int {
	if len(roman) == 0 {
		return 0
	}

	romanList := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	romanMap := map[string]int{
		"M":  1000,
		"CM": 900,
		"D":  500,
		"CD": 400,
		"C":  100,
		"XC": 90,
		"L":  50,
		"XL": 40,
		"X":  10,
		"IX": 9,
		"V":  5,
		"IV": 4,
		"I":  1,
	}

	num := 0

	for len(roman) > 0 {
		for _, r := range romanList {
			if strings.HasPrefix(roman, r) {
				num += romanMap[r]
				roman = strings.TrimPrefix(roman, r)

				break
			}
		}
	}

	return num
}
