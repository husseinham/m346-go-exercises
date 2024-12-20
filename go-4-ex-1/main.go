package main

import (
	"errors"
	"fmt"
)

func computeGrade(gotPoints, maxPoints float64) (float64, error) {
	if gotPoints < 0 || maxPoints <= 0 || gotPoints > maxPoints {
		return 0, errors.New("ungültige Punktzahlen")
	}
	return 1 + 5*(gotPoints/maxPoints), nil
}

func main() {
	grades := []struct {
		gotPoints float64
		maxPoints float64
	}{
		{17.5, 28.0},  // Erwartete Note: 4.125
		{28.0, 28.0},  // Erwartete Note: 6.0
		{0.0, 28.0},   // Erwartete Note: 1.0
		{30.0, 28.0},  // Fehlerfall
	}

	for _, g := range grades {
		grade, err := computeGrade(g.gotPoints, g.maxPoints)
		if err != nil {
			fmt.Printf("Error für Punkte %v/%v: %v\n", g.gotPoints, g.maxPoints, err)
		} else {
			fmt.Printf("Note für Punkte %v/%v: %.3f\n", g.gotPoints, g.maxPoints, grade)
		}
	}
}
