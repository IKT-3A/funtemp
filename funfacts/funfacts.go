package funfacts

type FunFacts struct {
	Sun   []string
	Luna  []string
	Terra []string
}

func GetFunFacts(about string) []string {
	facts := FunFacts{
		Sun: []string{
			"Temperature in the Sun's core is 15,000,000 degrees Celsius",
			"Temperature of the outer layer of the Sun is 5778 Kelvin",
		},
		Luna: []string{
			"Temperature of the Moon's surface at night is -183 degrees Celsius",
			"Temperature of the Moon's surface during the day is 106 degrees Celsius",
		},
		Terra: []string{
			"Highest temperature measured on the Earth's surface is 134 degrees Fahrenheit, 56.7 degrees Celsius, 329.82 Kelvin",
			"Lowest temperature measured on the Earth's surface is -89.4 degrees Celsius",
			"Temperature in the Earth's inner core is 9392 Kelvin",
		},
	}

	switch about {
	case "sun":
		return facts.Sun
	case "moon":
		return facts.Luna
	case "earth":
		return facts.Terra
	default:
		return []string{}
	}
}
