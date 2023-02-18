package space

const (
	orbitalEarthSeconds = float64(31557600) // 1.000.000.000
	orbitalMercurySeconds = 0.2408467 * orbitalEarthSeconds
	orbitalVenusSeconds = 0.61519726 * orbitalEarthSeconds
	orbitalMarsSeconds = 1.8808158 * orbitalEarthSeconds
	orbitalJupiterSeconds = 11.862615 * orbitalEarthSeconds
	orbitalSaturnSeconds = 29.447498 * orbitalEarthSeconds
	orbitalUranusSeconds = 84.016846 * orbitalEarthSeconds
	orbitalNeptuneSeconds = 164.79132 * orbitalEarthSeconds
)

type Planet string

const (
	Jupiter Planet = "Jupiter"
	Mars Planet = "Mars"
	Mercury Planet = "Mercury"
	Neptune Planet = "Neptune"
	Saturn Planet = "Saturn"
	Uranus Planet = "Uranus"
	Venus Planet = "Venus"
)

func Age(age float64, planet Planet) float64 {
	switch planet {
	case Mercury:
		return age/orbitalMercurySeconds
	case Venus:
		return age/orbitalVenusSeconds
	case Mars:
		return age/orbitalMarsSeconds
	case Jupiter:
		return age/orbitalJupiterSeconds
	case Saturn:
		return age/orbitalSaturnSeconds
	case Uranus:
		return age/orbitalUranusSeconds
	case Neptune:
		return age/orbitalNeptuneSeconds
	default:
		return age/orbitalEarthSeconds
	}
}

// best solution
/*
type Planet string

const Earth = 31557600.0

var PlanetSeconds = map[Planet]float64{
	"Earth":   Earth,
	"Mercury": Earth * 0.2408467,
	"Venus":   Earth * 0.61519726,
	"Mars":    Earth * 1.8808158,
	"Jupiter": Earth * 11.862615,
	"Saturn":  Earth * 29.447498,
	"Uranus":  Earth * 84.016846,
	"Neptune": Earth * 164.79132}

// Age accepts age in seconds and a planet name and returns a person's age on that planet
func Age(secs float64, planet Planet) float64 {
	return secs / PlanetSeconds[planet]

}
 */