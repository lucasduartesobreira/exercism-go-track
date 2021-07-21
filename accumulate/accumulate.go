// Package accumulate implement a method to apply operations over a collection of strings
package accumulate

// Accumulate take a collection of string and a operation to perform.
// It perform the operation over every element of the given collection
func Accumulate(collection []string, operation func(string) string) []string {
	collectionLength := len(collection)

	result := make([]string, 0, collectionLength)

	for _, element := range collection {
		result = append(result, operation(element))
	}

	return result
}
