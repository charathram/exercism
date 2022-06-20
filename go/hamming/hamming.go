// Utilities for DNA analysis.
package hamming

import (
	"errors"
)

// A function that calculates the hamming distance between
// two equal-length DNA strands, or returns an error if the
// strands are of unequal length.
func Distance(a, b string) (int, error) {
	var distance int
	if len(a) != len(b) {
		return -1, errors.New("the two strings are of unequal length")
	} else {
		ra := []rune(a)
		rb := []rune(b)
		for k, v := range ra {
			if rb[k] != v {
				distance += 1
			}
		}
	}
	return distance, nil
}
