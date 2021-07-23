// Package scrabble implement a method to find the scrabble score
package scrabble

import (
	"strings"
)

// letterToNumber maps each rune to its scrabble score
var letterToNumber = map[rune]int{
	'A': 1,
	'E': 1,
	'I': 1,
	'O': 1,
	'U': 1,
	'L': 1,
	'N': 1,
	'R': 1,
	'S': 1,
	'T': 1,
	'D': 2,
	'G': 2,
	'B': 3,
	'C': 3,
	'M': 3,
	'P': 3,
	'F': 4,
	'H': 4,
	'V': 4,
	'W': 4,
	'Y': 4,
	'K': 5,
	'J': 8,
	'X': 8,
	'Q': 10,
	'Z': 10,
}

// Score takes a string as input and return its Scrabble score
func Score(input string) int {
	sum := 0
	inputUpperCase := strings.ToUpper(input)

	strings.Map(func(r rune) rune {
		sum += letterToNumber[r]
		return r
	}, inputUpperCase)

	return sum
}
