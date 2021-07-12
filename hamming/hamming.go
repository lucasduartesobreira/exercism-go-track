package hamming

import (
	"errors"
)

//ErrorMessage is the error message for when both strings has different length
const ErrorMessage string = "Cannot compare two strings that differentiate on length"

// Distance receive strings a and b as parameters and return how many runes that differentiate the variable a from b.
// It can return a error if a and b has different length
func Distance(a, b string) (int, error) {
	if !hasSameLength(a, b) {
		return 0, errors.New(ErrorMessage)
	}
	return difference(a, b), nil
}

//hasSameLength return if strings a and b have the same length
func hasSameLength(a, b string) bool {
	return len(a) == len(b)
}

//difference return how many runes that differentiate the strings a from b
func difference(a, b string) int {
	hammingDistance := 0
	for i, actualRuneA := range a {
		actualRuneB := rune(b[i])
		if !isEqual(actualRuneA, actualRuneB) {
			hammingDistance++
		}
	}
	return hammingDistance
}

//isEqual receive runes a and b as parameters and return if they are equal
func isEqual(a rune, b rune) bool {
	return a == b
}
