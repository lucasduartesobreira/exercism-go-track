// Package acronym implement a method to get the abbreviation for any string
package acronym

import (
	"fmt"
	"strings"
)

// Abbreviate takes a string s and return its abbreviation
func Abbreviate(s string) string {

	filterSlash := strings.ReplaceAll(s, "-", " ")
	filterUnderline := strings.ReplaceAll(filterSlash, "_", " ")
	spaceSplitted := strings.Split(filterUnderline, " ")

	result := ""
	for _, str := range spaceSplitted {
		if len(str) != 0 {
			result = fmt.Sprintf("%s%c", result, (strings.ToUpper(str))[0])
		}
	}

	return result
}
