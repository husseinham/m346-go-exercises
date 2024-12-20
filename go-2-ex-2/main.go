package main

import "fmt"

fibs := []int{1, 1}

	// Berechnung der nächsten drei Werte
	for i := 2; i < 5; i++ {
		fibs = append(fibs, fibs[i-1]+fibs[i-2])
	}

	// Berechnung von vier weiteren Werten und Anhängen
	for i := 0; i < 4; i++ {
		next := fibs[len(fibs)-1] + fibs[len(fibs)-2]
		fibs = append(fibs, next)
	}

	fmt.Printf("Fibonacci-Zahlen: %v\n", fibs)