package main

import (
	"flag"
	"fmt"

	"funtemps/conv"
	"funtemps/funfacts"
)

var (
	fahr    float64
	celsius float64
	kelvin  float64
	out     string
	t       string
)

func init() {
	flag.Float64Var(&fahr, "F", 0.0, "temperature in degrees Fahrenheit")
	flag.Float64Var(&celsius, "C", 0.0, "temperature in degrees Celsius")
	flag.Float64Var(&kelvin, "K", 0.0, "temperature in Kelvin")
	flag.StringVar(&out, "out", "C", "output temperature scale: C - Celsius, F - Fahrenheit, K - Kelvin")
	flag.StringVar(&t, "t", "C", "temperature scale for fun facts: C - Celsius, F - Fahrenheit, K - Kelvin")
}

func main() {
	flag.Parse()

	// Check that only one of the temperature flags is set
	numTempFlags := 0
	if fahr != 0 {
		numTempFlags++
	}
	if celsius != 0 {
		numTempFlags++
	}
	if kelvin != 0 {
		numTempFlags++
	}
	if numTempFlags != 1 {
		fmt.Println("Please specify exactly one of the temperature flags (-F, -C, -K)")
		return
	}

	// Convert input temperature to chosen output scale
	var converted float64
	switch {
	case fahr != 0:
		switch out {
		case "C":
			converted = conv.FahrenheitToCelsius(fahr)
		case "K":
			converted = conv.FahrenheitToKelvin(fahr)
		case "F":
			converted = fahr
		}
	case celsius != 0:
		switch out {
		case "F":
			converted = conv.CelsiusToFahrenheit(celsius)
		case "K":
			converted = conv.CelsiusToKelvin(celsius)
		case "C":
			converted = celsius
		}
	case kelvin != 0:
		switch out {
		case "F":
			converted = conv.KelvinToFahrenheit(kelvin)
		case "C":
			converted = conv.KelvinToCelsius(kelvin)
		case "K":
			converted = kelvin
		}
	}

	// Get fun facts if -funfacts flag is used
	if t != "C" && (funfacts.GetFunFacts()[] == "sun" || funfacts == "moon" || funfacts == "earth") {
		fmt.Println("Fun facts are only available in Celsius")
		return
	}
	if funfacts != "sun" && funfacts != "moon" && funfacts != "earth" {
		fmt.Println("Invalid value for -funfacts. Please choose sun, moon, or earth.")
		return
	}
	if funfacts == "sun" {
		sunFacts := funfacts.GetFunFacts("sun")
		for _, fact := range sunFacts {
			fmt.Println(fact)
		}
	} else if funfacts == "moon" {
		moonFacts := funfacts.GetFunFacts("moon")
		if out != "C" {
			fmt.Println("Moon facts are only available in Celsius")
			return
		}
		for _, fact := range moonFacts {
			fmt.Println(fact)
		}
	} else if funfacts == "earth" {
		earthFacts := funfacts.GetFunFacts("earth")
		if out != "C" {
			for _, fact := range earthFacts {
			fmt.Println(fact)
			}
			}
			}

			if funfacts == "sun" || funfacts == "moon" || funfacts == "earth" {
				return
			}
			
			// Output the converted temperature
			switch out {
			case "C":
				fmt.Printf("Output: %.2fK er %.2f°C\n", converted, Celsius)
			case "F":
				fmt.Printf("Output: %.2fK er %.2f°F\n", converted, fahr)
			case "K":
				fmt.Printf("Output: %.2fK er %.2fK\n", converted, kelvin)
			}
			
		}

		// The function checks if the flag is specified on the command line
		// You don't have to use it, but it can help with the logic
		func isFlagPassed(name string) bool {
		found := false
		flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
		found = true
		}
		})
		return found
		}