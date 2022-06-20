// A package with a single function to share something with you while also sharing with me.
package twofer

import "fmt"

// ShareWith returns a string that says
// One for <name>, one for me or
// One for you, one for me if the name is empty.
func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	} else {
		return fmt.Sprintf("One for %s, one for me.", name)
	}
}
