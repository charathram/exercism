// A package that calculates the scrabble score
package scrabble

import "strings"

// A function that calculates the score of a Scrabble play given an input string.
func Score(word string) int {
	var score int = 0

	if word == "" {
		return 0
	}
	// Initialize a look up table to map letters with their points.
	valMap := makeMap()
	// convert the string to lowercase so that we don't run
	// into any case sensitivity issues.
	word = strings.ToLower(word)
	for _, letter := range word {
		score += valMap[string(letter)]
	}
	return score
}

// a private
func makeMap() map[string]int {
	m := make(map[string]int)

	m["a"] = 1
	m["b"] = 3
	m["c"] = 3
	m["d"] = 2
	m["e"] = 1
	m["f"] = 4
	m["g"] = 2
	m["h"] = 4
	m["i"] = 1
	m["j"] = 8
	m["k"] = 5
	m["l"] = 1
	m["m"] = 3
	m["n"] = 1
	m["o"] = 1
	m["p"] = 3
	m["q"] = 10
	m["r"] = 1
	m["s"] = 1
	m["t"] = 1
	m["u"] = 1
	m["v"] = 4
	m["w"] = 4
	m["x"] = 8
	m["y"] = 4
	m["z"] = 10

	return m
}
