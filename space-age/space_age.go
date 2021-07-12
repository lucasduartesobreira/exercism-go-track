package space

type Planet string

const (
	Mercury Planet = "Mercury"
	Venus   Planet = "Venus"
	Earth   Planet = "Earth"
	Mars    Planet = "Mars"
	Jupiter Planet = "Jupiter"
	Saturn  Planet = "Saturn"
	Uranus  Planet = "Uranus"
	Neptune Planet = "Neptune"
)

var YearInEarthPerYearInPlanet = map[Planet]float64{
	Mercury: 0.2408467,
	Venus:   0.61519726,
	Earth:   1.0,
	Mars:    1.8808158,
	Jupiter: 11.862615,
	Saturn:  29.447498,
	Uranus:  84.016846,
	Neptune: 164.79132,
}

const SecondsPerEarthYear float64 = 31557600

func Age(seconds float64, planet Planet) float64 {
	earthYears := seconds / SecondsPerEarthYear
	yearsOnPlanet := earthYears / YearInEarthPerYearInPlanet[planet]
	return yearsOnPlanet
}
