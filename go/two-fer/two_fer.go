// Package twofer provides methods to share between me and other person
package twofer

import "fmt"

/*
ShareWith receives an optional param #{name} and returns
"One for %s, one for me." replacing %s with #{name}
value default value for is name="you"
*/
func ShareWith(name string) string {
	// check if name has no len()
	if len(name) == 0 {
		// default value if empty
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
