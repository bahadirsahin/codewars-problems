package set5

import "strings"

func Q147(num int) string {
	if num == 0 || num >= 4000 {
		return ""
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

	res := ""

	for _, r := range romanList {
		d := num / romanMap[r]

		if d == 0 {
			continue
		}

		num -= d * romanMap[r]
		res += strings.Repeat(r, d)
	}

	return res
}
