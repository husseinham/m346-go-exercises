package main

import "fmt"

func main() {
    // Vorhandene Umrechnung (imperial zu metrisch)
    miles := 10.0
    fahrenheit := 68.0

    kilometers := miles * 1.60934
    celsius := (fahrenheit - 32) * 5 / 9

    fmt.Printf("%.2f Meilen sind %.2f Kilometer\n", miles, kilometers)
    fmt.Printf("%.2f °F sind %.2f °C\n", fahrenheit, celsius)

    // TODO: Umrechnung in die andere Richtung (metrisch zu imperial)
    kilometersToMiles := kilometers / 1.60934
    celsiusToFahrenheit := celsius*9/5 + 32

    fmt.Printf("%.2f Kilometer sind %.2f Meilen\n", kilometers, kilometersToMiles)
    fmt.Printf("%.2f °C sind %.2f °F\n", celsius, celsiusToFahrenheit)
}
