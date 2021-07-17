// Package raindrops implement custom messages based on rain drops
package raindrops

import "strconv"

// factorAndMessage is the map between the factor and its specified message
var factorAndMessage = map[int]string{
	3: "Pling",
	5: "Plang",
	7: "Plong",
}

// Convert take raindrops as parameter and return its custom message
// depending of raindrops factors.
// If raindrops has 3 as a factor, will add 'Pling' to the result
// If raindrops has 5 as a factor, will add 'Plang' to the result
// If raindrops has 7 as a factor, will add 'Plong' to the result
// If raindrops does not have any of 3,5 or 7 as a factor, return raindrops
func Convert(raindrops int) string {
	result := makeResponse(raindrops)
	return result
}

// makeResponse take raindrops as parameter and will make the custom message
// depending on raindrops and given the rules specified.
func makeResponse(raindrops int) string {
	result := ""

	for factor, message := range factorAndMessage {
		if isFactor(raindrops, factor) {
			result += message
		}
	}

	if result == "" {
		return strconv.Itoa(raindrops)
	}
	return result
}

// isFactor verify if toTest has factor as a factor
func isFactor(toTest, factor int) bool {
	return toTest%factor == 0
}
