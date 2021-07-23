// Package collatzconjecture implements a method to get the number of
// steps required by the Collatz Conjecture
package collatzconjecture

import "errors"

// CollatzConjecture takes an int as a parameter and returns how many steps were required to
// get to 1 following the Collatz Conjecture
func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("Not able to calculate the number of steps for n <= 0")
	}

	count := 0
	dummyVar := n

	for dummyVar > 1 {
		if dummyVar%2 == 0 {
			dummyVar /= 2
		} else {
			dummyVar = 3*dummyVar + 1
		}
		count++
	}

	return count, nil
}
