// Package dna contains implementation of a DNA type
package dna

import (
	"errors"
)

type Histogram map[rune]int

type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (histogram Histogram, err error) {
	histogram = Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}

	for _, rune := range d {
		if _, ok := histogram[rune]; !ok {
			return histogram, errors.New("Invalid nucleotide")
		}
		histogram[rune] += 1
	}

	return histogram, nil
}
