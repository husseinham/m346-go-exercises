package main

import "fmt"

func main() {
	// TODO: Variablen deklarieren und initialisieren
	var firstName string = "Max"
	var lastName string = "Muster"
	var age int = 25
	var dayOfBirth int = 15
	var monthOfBirth string = "Mai"
	var yearOfBirth int = 1998
	var zodiacSign rune = '\u264A' // Beispiel: Waage (Unicode: ♊)
	var hobbies []string = []string{"Programmieren", "Lesen", "Wandern"}
	var siblings int = 2
	var favoriteColor string = "Blau"

	// Ausgabe
	fmt.Println("Steckbrief")
	fmt.Println("Vorname:", firstName)
	fmt.Println("Nachname:", lastName)
	fmt.Println("Alter:", age)
	fmt.Println("Geburtstag:", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Println("Sternzeichen:", string(zodiacSign))
	fmt.Println("Hobbies:", hobbies)
	fmt.Println("Anzahl Geschwister:", siblings)
	fmt.Println("Lieblingsfarbe:", favoriteColor)

	// TODO: Anzahl der Geschwister um 1 erhöhen
	siblings++
	fmt.Println("Neue Anzahl Geschwister:", siblings)
}
