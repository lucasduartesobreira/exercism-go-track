package etl

import "strings"

func Transform(oldScrablle map[int][]string) (newScrabble map[string]int) {
	newScrabble = make(map[string]int, len(oldScrablle))
	for point, letters := range oldScrablle {
		for _, letter := range letters {
			newScrabble[strings.ToLower(letter)] = point
		}
	}
	return
}
