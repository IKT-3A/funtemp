package funfacts

import (
	"reflect"
	"testing"
)

func TestGetFunFacts(t *testing.T) {
	type test struct {
		input string
		want  []string
	}
	tests := []test{
		{input: "sun", want: []string{
			"The sun is a star, not a planet.",
			"The sun contains 99.86% of the total mass of the Solar System.",
			"The temperature in the Sun's core is 15 million degrees Celsius.",
		}},
		{input: "moon", want: []string{
			"The moon is Earth's only natural satellite.",
			"The moon is not a perfect sphere.",
			"The moon has no atmosphere.",
		}},
		{input: "earth", want: []string{
			"Earth is the third planet from the Sun.",
			"Earth is the only planet known to support life.",
			"The Earth is not a perfect sphere, it is an oblate spheroid.",
		}},
		{input: "mars", want: []string{}},
	}

	for _, tc := range tests {
		if got := GetFunFacts(tc.input); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("GetFunFacts(%v) = %v; want %v", tc.input, got, tc.want)
		}
	}
}
