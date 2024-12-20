package main

import "fmt"

// Funktionen
func convertCelsiusToFahrenheit(c float64) float64 {
	return c*9/5 + 32
}

func convertFahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}

// Typen und Methoden
type Celsius float64
type Fahrenheit float64

func (c Celsius) ConvertToFahrenheit() Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func (f Fahrenheit) ConvertToCelsius() Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func main() {
	// Testfälle für Funktionen
	fmt.Println("23°C -> Fahrenheit:", convertCelsiusToFahrenheit(23.0)) // Erwartet: 73.4
	fmt.Println("73.4°F -> Celsius:", convertFahrenheitToCelsius(73.4)) // Erwartet: 23.0

	// Testfälle für Methoden
	var cozy Celsius = 23.0
	fmt.Println("23°C -> Fahrenheit (Methode):", cozy.ConvertToFahrenheit()) // Erwartet: 73.4
	fmt.Println("73.4°F -> Celsius (Methode):", Fahrenheit(73.4).ConvertToCelsius()) // Erwartet: 23.0
}
