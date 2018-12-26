package space

var m = make(map[Planet]float64)

type Planet string

func init() {
	m["Mercury"] = 0.2408467
	m["Venus"] = 0.61519726
	m["Mars"] = 1.8808158
	m["Jupiter"] = 11.862615
	m["Saturn"] = 29.447498
	m["Neptune"] = 164.79132
	m["Uranus"] = 84.016846
}

const earthYearInSeconds = 31557600

//Age calculates your age in planet years given seconds
func Age(seconds float64, planet Planet) float64 {
	earthYears := seconds / earthYearInSeconds
	if planet == "Earth" {
		return float64(earthYears)
	}
	return float64(earthYears) / m[planet]
}
