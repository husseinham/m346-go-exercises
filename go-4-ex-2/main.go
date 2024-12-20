package main

import (
	"fmt"
	"math"
)

// Funktion
func computeHypotenuse(a, b float64) float64 {
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

// Zusatzaufgabe: Typ und Methode
type ShortSides struct {
	a, b float64
}

func (s ShortSides) Hypotenuse() float64 {
	return computeHypotenuse(s.a, s.b)
}

func main() {
	// Testfälle für computeHypotenuse
	fmt.Println("Hypotenuse (3, 4):", computeHypotenuse(3, 4)) // Erwartet: 5.0
	fmt.Println("Hypotenuse (5, 12):", computeHypotenuse(5, 12)) // Erwartet: 13.0
	fmt.Println("Hypotenuse (8, 15):", computeHypotenuse(8, 15)) // Erwartet: 17.0

	// Testfälle für ShortSides
	ss := ShortSides{a: 3, b: 4}
	fmt.Println("Hypotenuse via Methode:", ss.Hypotenuse()) // Erwartet: 5.0
}
