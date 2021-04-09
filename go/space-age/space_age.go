package space

type Planet string

func Age(seconds float64, planet Planet) float64 {
	var planetarySeconds float64
	secondsInOneEarthYear := float64(31557600)

	switch planet {
	case "Mercury":
		planetarySeconds = secondsInOneEarthYear * 0.2408467
	case "Venus":
		planetarySeconds = secondsInOneEarthYear * 0.61519726
	case "Earth":
		planetarySeconds = secondsInOneEarthYear
	case "Mars":
		planetarySeconds = secondsInOneEarthYear * 1.8808158
	case "Jupiter":
		planetarySeconds = secondsInOneEarthYear * 11.862615
	case "Saturn":
		planetarySeconds = secondsInOneEarthYear * 29.447498
	case "Uranus":
		planetarySeconds = secondsInOneEarthYear * 84.016846
	case "Neptune":
		planetarySeconds = secondsInOneEarthYear * 164.79132
	}

	planetAge := seconds / planetarySeconds
	return planetAge
}
