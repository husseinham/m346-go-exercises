package main

import "fmt"

type FullName struct {
	// TODO: add fields
}

// Lösung für die Aufgaben

// Aufgabe 1: Steckbrief II
package main

import "fmt"

type Person struct {
	FirstName     string
	LastName      string
	DayOfBirth    int
	MonthOfBirth  int
	YearOfBirth   int
	NumberOfSiblings int
}

func main() {
	// Initialisierung des Steckbriefs
	me := Person{
		FirstName:        "Hussein",
		LastName:         "Hammade",
		DayOfBirth:       15,
		MonthOfBirth:     7,
		YearOfBirth:      1995,
		NumberOfSiblings: 2,
	}

	// Anzahl der Geschwister erhöhen
	me.NumberOfSiblings++

	fmt.Printf("Mein Steckbrief: %+v\n", me)
}