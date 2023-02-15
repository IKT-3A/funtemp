package conv

func convert(value float64, inputUnit, outputUnit string) float64 {

	var kelvin float64
	switch inputUnit {
	case "K":
		kelvin = value
	case "C":
		kelvin = value + 273.15
	case "F":
		kelvin = (value + 459.67) * 5 / 9
	default:
		panic("Unknown input temperature unit")
	}

	var outputValue float64
	switch outputUnit {
	case "K":
		outputValue = kelvin
	case "C":
		outputValue = kelvin - 273.15
	case "F":
		outputValue = kelvin*9/5 - 459.67
	default:
		panic("Unknown output temperature unit")
	}

	return outputValue
}
