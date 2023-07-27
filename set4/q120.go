package set4

import "strings"

func Q120(morseCode string) string {
	mc := map[string]string{
		".-":        "A",
		"-...":      "B",
		"-.-.":      "C",
		"-..":       "D",
		".":         "E",
		"..-.":      "F",
		"--.":       "G",
		"....":      "H",
		"..":        "I",
		".---":      "J",
		"-.-":       "K",
		".-..":      "L",
		"--":        "M",
		"-.":        "N",
		"---":       "O",
		".--.":      "P",
		"--.-":      "Q",
		".-.":       "R",
		"...":       "S",
		"-":         "T",
		"..-":       "U",
		"...-":      "V",
		".--":       "W",
		"-..-":      "X",
		"-.--":      "Y",
		"--..":      "Z",
		".----":     "1",
		"..---":     "2",
		"...--":     "3",
		"....-":     "4",
		".....":     "5",
		"-....":     "6",
		"--...":     "7",
		"---..":     "8",
		"----.":     "9",
		"-----":     "0",
		"...---...": "SOS",
		"-.-.--":    "!",
		".-.-.-":    ".",
	}

	morseCode = strings.TrimSpace(morseCode)

	// base
	if len(morseCode) == 0 {
		return ""
	}

	// check for single rune / word
	if len(mc[morseCode]) > 0 {
		return mc[morseCode]
	}

	words := strings.Split(morseCode, "   ")
	res := ""

	for _, word := range words {
		letters := strings.Split(word, " ")

		for _, letter := range letters {
			res += mc[letter]
		}

		res += " "
	}

	res = strings.TrimSpace(res)

	return res
}
