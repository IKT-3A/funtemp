package funcfacts

import "testing"

func TestFunFacts(t *testing.T) {
	expected := map[string]map[string]interface{}{
		"sun": {
			"temperature": 5506.85,
			"unit":        "C",
		},
		"core of the sun": {
			"temperature": 15000000,
			"unit":        "C",
		},
		"boiling water": {
			"temperature": 100,
			"unit":        "C",
		},
		"freezing point of water": {
			"temperature": 0,
			"unit":        "C",
		},
		"highest recorded temperature on Earth": {
			"temperature": 56.7,
			"unit":        "C",
		},
		"lowest recorded temperature on Earth": {
			"temperature": -89.2,
			"unit":        "C",
		},
	}

	for name, fact := range funFacts {
		if _, ok := expected[name]; !ok {
			t.Errorf("Unknown fun fact: %s", name)
			continue
		}
		expectedFact := expected[name]
		if fact["temperature"] != expectedFact["temperature"] {
			t.Errorf("Unexpected temperature for %s: got %v, want %v", name, fact["temperature"], expectedFact["temperature"])
		}
		if fact["unit"] != expectedFact["unit"] {
			t.Errorf("Unexpected temperature unit for %s: got %v, want %v", name, fact["unit"], expectedFact["unit"])
		}
	}
}
