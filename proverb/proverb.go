// Package proverb implement methods to transform rhymes into the proverb
package proverb

import "fmt"

// Proverb take rhyme as parameter and return a slice containing all the lines
// of the proverb
func Proverb(rhyme []string) []string {
	result := generateProverb(rhyme)
	return result
}

// generateProverb take a rhyme as parameter and return each line of the proverb
// formatted as follows:
// - For want of a "rhyme" the "next rhyme" was lost. If isn't the last line
// - And all for the want of a "first rhyme". If is the last line
// If rhyme is empty then will return rhyme itself
func generateProverb(rhyme []string) []string {
	if isEmpty(rhyme) {
		return rhyme
	}

	rhymeLength := len(rhyme)
	result := make([]string, 0, rhymeLength+1)

	for index, actual := range rhyme {
		nextIndex := index + 1
		if isInsideBoundaries(nextIndex, rhymeLength) {
			next := rhyme[nextIndex]
			result = append(result, makeMiddlePhrase(actual, next))
		}
	}

	result = append(result, makeLastPhrase(rhyme[0]))

	return result
}

// makeMiddlePhrase return a string formatted following the middle phrase pattern
func makeMiddlePhrase(actual, next string) string {
	return fmt.Sprintf("For want of a %s the %s was lost.", actual, next)
}

// makeLastPhrase return a string formatted following the last phrase pattern
func makeLastPhrase(first string) string {
	return fmt.Sprintf("And all for the want of a %s.", first)
}

// isInsideBoundaries return if index is inside length boundaries, it's index < length
func isInsideBoundaries(index, length int) bool {
	return index < length
}

// isEmpty return if a string array is empty
func isEmpty(rhyme []string) bool {
	return len(rhyme) == 0
}
