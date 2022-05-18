package isogram

import "unicode"

func IsIsogram(word string) bool {
	set := make(map[rune]struct{}, 0)

	for _, ch := range word {
		lowercase := unicode.ToLower(ch)

		if !unicode.IsLetter(lowercase) {
			continue
		}

		if _, exists := set[lowercase]; exists {
			return false
		} else {
			set[lowercase] = struct{}{}
		}
	}

	return true
}
