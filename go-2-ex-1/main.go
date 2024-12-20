package main

import "fmt"

// Definieren der Datenstruktur
type Person struct {
	FirstName   string
	LastName    string
	DayOfBirth  int
	MonthOfBirth string
	YearOfBirth int
	Siblings    int
}

func main() {
	// Persönlicher Steckbrief
	me := Person{
		FirstName:   "Max",
		LastName:    "Muster",
		DayOfBirth:  15,
		MonthOfBirth: "Mai",
		YearOfBirth: 1990,
		Siblings:    2,
	}

	fmt.Println("Vor Anpassung:", me)

	// Anzahl der Geschwister erhöhen
	me.Siblings++
	fmt.Println("Nach Anpassung:", me)
}
