// Make the noise of raindrops depending on the input
package raindrops

import "fmt"

// Converts the input int to a string depending
// on whether the input is divisible by 3, 5, or 7.
func Convert(number int) string {
	var s string

	if number%3 == 0 {
		s += "Pling"
	}
	if number%5 == 0 {
		s += "Plang"
	}
	if number%7 == 0 {
		s += "Plong"
	}

	if s == "" {
		return fmt.Sprint(number)
	} else {
		return s
	}
}
