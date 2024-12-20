package main

import "fmt"

func main() {
	// Slice initialisieren
	fibs := make([]int, 5)
	fibs[0], fibs[1] = 1, 1

	// Berechnen der ersten fünf Fibonacci-Zahlen
	for i := 2; i < len(fibs); i++ {
		fibs[i] = fibs[i-1] + fibs[i-2]
	}
	fmt.Println("Erste fünf Fibonacci-Zahlen:", fibs)

	// Erweiterung des Slices um vier weitere Zahlen
	for i := 0; i < 4; i++ {
		next := fibs[len(fibs)-1] + fibs[len(fibs)-2]
		fibs = append(fibs, next)
	}
	fmt.Println("Erweiterte Fibonacci-Zahlen:", fibs)
}
