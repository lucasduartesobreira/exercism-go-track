// Package triangle implement methods to determine which kind of triangle the given lengths are.
package triangle

import "math"

type Kind int

const (
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

// KindFromSides return the kind of triangle the lengths make.
func KindFromSides(a, b, c float64) Kind {
	if isNotATriangle(a, b, c) {
		return NaT
	}

	if a == b && b == c {
		return Equ
	} else if a == b || a == c || b == c {
		return Iso
	}

	return Sca
}

// isNotATriangle return if the given lengths make a triangle.
func isNotATriangle(a, b, c float64) bool {
	return isInvalidLength(a, b, c) || a > b+c || b > a+c || c > a+b
}

// isInvalidLength return if the given lengths are invalid. That means a,b and c contained on the interval ]0,Inf[.
func isInvalidLength(a, b, c float64) bool {
	return math.IsInf(a+b+c, 0) || math.IsNaN(a+b+c) || a <= 0 || b <= 0 || c <= 0
}
