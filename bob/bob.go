// Package bob implements a method to make Bob's response
package bob

import (
	"strings"
)

// Hey returns Bob's response given the conversation
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	isQuestion := isQuestion(remark)
	isYelling := isYelling(remark)
	isEmpty := remark == ""

	switch {
	case isEmpty:
		return "Fine. Be that way!"
	case isQuestion && isYelling:
		return "Calm down, I know what I'm doing!"
	case isYelling:
		return "Whoa, chill out!"
	case isQuestion:
		return "Sure."
	default:
		return "Whatever."
	}
}

// isQuestion returns if the inputs is a question
func isQuestion(remark string) bool {
	return strings.HasSuffix(remark, "?")
}

// isYelling returns if the inputs is in all capitals
func isYelling(remark string) bool {
	containsOnlyNotLetters := true

	for _, v := range remark {
		if v >= 'a' && v <= 'z' {
			return false
		}
		if v >= 'A' && v <= 'Z' {
			containsOnlyNotLetters = false
		}
	}

	if containsOnlyNotLetters {
		return false
	}

	return true
}
