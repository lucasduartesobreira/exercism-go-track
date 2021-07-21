// Package strand implement methods to transcribe DNA strand into RNA strand
package strand

import "strings"

// ToRNA takes a DNA strand and return its transcribed RNA strand
func ToRNA(dna string) string {
	return strings.Map(rnaComplement, dna)
}

// rnaComplement takes a DNA nucleotide and return the RNA complement
func rnaComplement(dna rune) rune {
	switch dna {
	case 'G':
		return 'C'
	case 'C':
		return 'G'
	case 'T':
		return 'A'
	default:
		return 'U'
	}
}
