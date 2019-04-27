// Package twofer implements ShareWith function
package twofer

// ShareWith takes a name as an argument and returns "One for <name>, one for me."
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
