package main

import (
	"fmt"
	"math"
)

// Funktion für Diskriminante
func computeDiscriminant(a, b, c float64) float64 {
	return math.Pow(b, 2) - 4*a*c
}

// Funktion für Lösungen
func computeQuadraticFormula(a, b, c float64) []float64 {
	D := computeDiscriminant(a, b, c)
	if D < 0 {
		return []float64{}
	} else if D == 0 {
		return []float64{-b / (2 * a)}
	} else {
		return []float64{
			(-b + math.Sqrt(D)) / (2 * a),
			(-b - math.Sqrt(D)) / (2 * a),
		}
	}
}

func main() {
	// Testfälle
	fmt.Println("Lösungen für D > 0 (3, 4, 1):", computeQuadraticFormula(3, 4, 1)) // Erwartet: [-0.3333, -1.0]
	fmt.Println("Lösungen für D = 0 (2, 4, 2):", computeQuadraticFormula(2, 4, 2)) // Erwartet: [-1.0]
	fmt.Println("Lösungen für D < 0 (3, 4, 2):", computeQuadraticFormula(3, 4, 2)) // Erwartet: []
}
