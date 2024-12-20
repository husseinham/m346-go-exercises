package main

import "fmt"

// Datenstrukturen definieren
type Student struct {
	FirstName string
	LastName  string
}

type Class struct {
	Name     string
	Students []Student
}

type Module struct {
	Number int
	Classes []Class
}

func main() {
	// Klassen erstellen
	classA := Class{
		Name: "Class A",
		Students: []Student{
			{"Anna", "Müller"},
			{"Peter", "Schmidt"},
			{"Lisa", "Weber"},
		},
	}
	classB := Class{
		Name: "Class B",
		Students: []Student{
			{"Tom", "Hoffmann"},
			{"Sara", "Klein"},
			{"Mark", "Fischer"},
		},
	}

	// Module erstellen
	modules := []Module{
		{
			Number: 346,
			Classes: []Class{classA, classB},
		},
		{
			Number: 350,
			Classes: []Class{classA},
		},
		{
			Number: 360,
			Classes: []Class{classB},
		},
	}

	// Daten ausgeben
	for _, module := range modules {
		fmt.Printf("Modul %d wird von folgenden Klassen besucht:\n", module.Number)
		for _, class := range module.Classes {
			fmt.Printf("  %s mit den Schülern:\n", class.Name)
			for _, student := range class.Students {
				fmt.Printf("    %s %s\n", student.FirstName, student.LastName)
			}
		}
	}
}
