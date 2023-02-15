package conv

import "testing"

func TestConvert(t *testing.T) {
	tests := []struct {
		value    float64
		fromUnit string
		toUnit   string
		expected float64
	}{
		{0, "K", "C", -273.15},
		{100, "C", "F", 212},
		{37, "C", "K", 310.15},
		{1, "F", "C", -17.2222},
		{1, "F", "K", 255.928},
	}

	for _, test := range tests {
		result := convert(test.value, test.fromUnit, test.toUnit)
		if result != test.expected {
			t.Errorf("convert(%g, %s, %s) = %g; want %g", test.value, test.fromUnit, test.toUnit, result, test.expected)
		}
	}
}
