package conv

// Converts Fahrenheit to Celsius
func FahrenheitToCelsius(value float64) float64 {
	return (value - 32) * 5 / 9
}

// Converts Celsius to Fahrenheit
func CelsiusToFahrenheit(value float64) float64 {
	return (value * 9 / 5) + 32
}

// Converts Kelvin to Fahrenheit
func KelvinToFahrenheit(value float64) float64 {
	return (value * 9 / 5) - 459.67
}

// Converts Fahrenheit to Kelvin
func FahrenheitToKelvin(value float64) float64 {
	return (value + 459.67) * 5 / 9
}

// Converts Celsius to Kelvin
func CelsiusToKelvin(value float64) float64 {
	return value + 273.15
}

// Converts Kelvin to Celsius
func KelvinToCelsius(value float64) float64 {
	return value - 273.15
}
