package main

import (
	"flag"
	"fmt"
	"funtemps/conv"
	funcfacts "funtemps/funfacts"
)

var value float64
var toScale string
var funFact string

func init() {
	flag.Float64Var(&value, "value", 0.0, "Temperature value")
	flag.StringVar(&toScale, "to", "C", "Convert to scale (C, F, K)")
	flag.StringVar(&funFact, "fact", "", "Fun fact subject (sun, moon, earth)")
}

func main() {
	flag.Parse()

	if funFact != "" {
		if toScale != "C" {
			fmt.Println("Error: Fun facts can only be shown in Celsius (C) scale.")
			return
		}
		facts := funcfacts.GetFunFacts(funFact)
		for _, fact := range facts {
			fmt.Println(fact)
		}
		return
	}

	switch toScale {
	case "C":
		celsius := value
		if toScale == "F" {
			celsius = conv.FahrenheitToCelsius(value)
		} else if toScale == "K" {
			celsius = conv.KelvinToCelsius(value)
		}
		fmt.Printf("Output: %.2f%s is %.2f°C\n", value, toScale, celsius)
	case "F":
		fahrenheit := value
		if toScale == "C" {
			fahrenheit = conv.CelsiusToFahrenheit(value)

			if toScale == "K" {
				fahrenheit = conv.KelvinToFahrenheit(value)
			}
			fmt.Printf("Output: %.2f%s is %.2f°F\n", value, toScale, fahrenheit)
		}
	case "K":
		kelvin := value
		if toScale == "C" {
			kelvin = conv.CelsiusToKelvin(value)
		} else if toScale == "F" {
			kelvin = conv.FahrenheitToKelvin(value)
		}
		fmt.Printf("Output: %.2f%s is %.2fK\n", value, toScale, kelvin)
	default:
		fmt.Println("Error: Invalid scale specified. Valid options are C, F, and K.")
	}
}
