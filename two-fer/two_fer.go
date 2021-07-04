package twofer

import "fmt"

// ShareWith will take a name as parameter and return "One for X, one for me."
// where X is the name. If the name is missing, it will return "One for you, one for me."
func ShareWith(name string) string {
	normalizedName := normalizeName(name)

	return fmt.Sprintf("One for %s, one for me.", normalizedName)
}

//normalizeName will take name as parameter and return either "you", if name is missing
//or name, otherwise
func normalizeName(name string) string {
	if name == "" {
		return "you"
	}
	return name
}
