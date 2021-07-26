package darts

import (
	"math"
)

type Area int

const (
	outsideCircle = 0
	outerCircle   = 1
	middleCircle  = 5
	innerCircle   = 10
)

var areaGivenLimits = map[struct {
	outsideRadius float64
	insideRadius  float64
}]Area{
	{10.0, 5.0}: outerCircle,
	{5.0, 1.0}:  middleCircle,
	{1.0, 0.0}:  innerCircle,
}

func Score(x, y float64) int {
	area := getAreaOfLanding(x, y)

	return int(area)
}

func getAreaOfLanding(x, y float64) Area {
	if x == 0 && y == 0 {
		return innerCircle
	}

	distance := getDistanceFromCenter(x, y)

	for k, v := range areaGivenLimits {
		if distance > k.insideRadius && distance <= k.outsideRadius {
			return v
		}
	}

	return outsideCircle
}

func getDistanceFromCenter(x, y float64) float64 {
	return math.Abs(math.Sqrt(x*x + y*y))
}
