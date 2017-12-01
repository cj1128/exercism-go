package space

type Planet = string

const earthSeconds float64 = 31557600

var planetSeconds = map[string]float64{
	"Earth":   earthSeconds,
	"Mercury": 0.2408467 * earthSeconds,
	"Venus":   0.61519726 * earthSeconds,
	"Mars":    1.8808158 * earthSeconds,
	"Jupiter": 11.862615 * earthSeconds,
	"Saturn":  29.447498 * earthSeconds,
	"Uranus":  84.016846 * earthSeconds,
	"Neptune": 164.79132 * earthSeconds,
}

func Age(seconds float64, planet Planet) float64 {
	return seconds / planetSeconds[planet]
}
