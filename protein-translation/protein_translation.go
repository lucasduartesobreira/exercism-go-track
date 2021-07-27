package protein

import (
	"errors"
)

var ErrStop error = errors.New("Err(Stop): Translation ended, found a STOP codons")
var ErrInvalidBase error = errors.New("Err(InvalidBase): Codon has a invalid base")

func FromCodon(codon string) (protein string, err error) {
	switch codon {
	case "AUG":
		protein = "Methionine"
		err = nil
		break

	case "UUU", "UUC":
		protein = "Phenylalanine"
		err = nil
		break

	case "UUA", "UUG":
		protein = "Leucine"
		err = nil
		break

	case "UCU", "UCC", "UCA", "UCG":
		protein = "Serine"
		err = nil
		break

	case "UAU", "UAC":
		protein = "Tyrosine"
		err = nil
		break

	case "UGU", "UGC":
		protein = "Cysteine"
		err = nil
		break

	case "UGG":
		protein = "Tryptophan"
		err = nil
		break

	case "UAA", "UAG", "UGA":
		protein = ""
		err = ErrStop
		break

	default:
		protein = ""
		err = ErrInvalidBase
		break
	}

	return
}

func FromRNA(rnaSequence string) (proteins []string, err error) {
	proteins = make([]string, 0, len(rnaSequence)/3)

	rnaIterator := New(rnaSequence)

	for codon := rnaIterator.next(); codon != ""; codon = rnaIterator.next() {
		protein, err := FromCodon(codon)

		if err == ErrStop {
			return proteins, nil
		}

		if err != nil {
			return proteins, err
		}

		proteins = append(proteins, protein)
	}

	return
}

type Iter struct {
	string
	length int
}

func New(str string) *Iter {
	return &Iter{string: str, length: len(str)}
}

func (m *Iter) next() (next string) {
	if m.length >= 2 {
		next = m.string[0:3]
		m.string = m.string[3:]
		m.length -= 3
	}

	return
}
