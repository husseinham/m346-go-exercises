package main

import "fmt"

func main() {
	// Map erstellen
	modules := map[int]string{
		346: "Cloud-Lösungen konzipieren und realisieren",
		350: "Datenbanklösungen implementieren",
		360: "Sicherheit in IT-Systemen gewährleisten",
	}

	// Ausgabe der Modulinformationen
	fmt.Println("Modul 346:", modules[346])
	fmt.Println("Modul 350:", modules[350])
	fmt.Println("Modul 360:", modules[360])

	// Manipulationen an der Map
	delete(modules, 350) // Element entfernen
	modules[370] = "Webentwicklung umsetzen" // Element hinzufügen
	modules[346] = "Cloud-Lösungen modernisieren" // Element ersetzen

	// Map nach Anpassungen ausgeben
	fmt.Println("Aktualisierte Module:", modules)
}
