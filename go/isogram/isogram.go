package isogram

import "strings"

func IsIsogram(word string) bool {
	letterCount := make(map[string]int)
	var isIsogram bool = true

	if word == "" {
		return true
	}

	word = strings.ToLower(word)

	for _, letter := range word {
		_, exists := letterCount[string(letter)]
		// if the letter exists in the map, it means
		// this is the second occurrence of the letter
		// and so, it's not an isogram.
		if exists {
			if (letter != ' ') && (letter != '-') {
				isIsogram = false
			}
		} else {
			letterCount[string(letter)] = 1
		}
	}
	return isIsogram
}
