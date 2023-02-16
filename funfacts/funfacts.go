package funfacts

// Structure for funfacts
type FunFacts struct {
	Sun   []string
	Luna  []string
	Terra []string
}

// Returns funfacts for the given subject
func GetFunFacts(subject string) []string {
	facts := FunFacts{
		Sun: []string{
			"The Sun is a nearly perfect sphere.",
			"The Sun is about 109 times larger than Earth.",
		},
		Luna: []string{
			"The Moon is a natural satellite of the Earth.",
			"The Moon is about one quarter the size of the Earth.",
		},
		Terra: []string{
			"Earth is the only planet known to support life.",
			"Earth is the third planet from the Sun.",
		},
	}
	switch subject {
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
