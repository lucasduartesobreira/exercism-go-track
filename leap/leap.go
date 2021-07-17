// Package leap implement methods to verify if a given year is a leap year.
package leap

// IsLeapYear take year as parameter and return if its a leap year.
func IsLeapYear(year int) bool {
	if isDivisableBy(year, 400) {
		return true
	} else if isDivisableBy(year, 4) && !isDivisableBy(year, 100) {
		return true
	}
	return false
}

// Return if year is divisible by dividend.
func isDivisableBy(year, dividend int) bool {
	return year%dividend == 0
}
