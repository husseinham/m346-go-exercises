package main

func main() {
	type Student struct {
		FirstName string
		LastName  string
	}

	type Class struct {
		ClassName string
		Students  []Student
	}

	type Module struct {
		ModuleID   int
		Classes    []Class
	}

	// Beispieldaten
	classA := Class{
		ClassName: "Klasse A",
		Students: []Student{
			{FirstName: "Max", LastName: "Mustermann"},
			{FirstName: "Erika", LastName: "Musterfrau"},
			{FirstName: "John", LastName: "Doe"},
		},
	}

	classB := Class{
		ClassName: "Klasse B",
		Students: []Student{
			{FirstName: "Jane", LastName: "Doe"},
			{FirstName: "Anna", LastName: "Smith"},
			{FirstName: "Peter", LastName: "Parker"},
		},
	}

	module1 := Module{
		ModuleID: 346,
		Classes:  []Class{classA, classB},
	}

	module2 := Module{
		ModuleID: 123,
		Classes:  []Class{classA},
	}

	module3 := Module{
		ModuleID: 890,
		Classes:  []Class{classB},
	}

	// Ausgabe der Daten
	modulesList := []Module{module1, module2, module3}
	fmt.Println("Module und zugehörige Klassen:")
	for _, module := range modulesList {
		fmt.Printf("Modul %d:\n", module.ModuleID)
		for _, class := range module.Classes {
			fmt.Printf("  Klasse: %s\n", class.ClassName)
			for _, student := range class.Students {
				fmt.Printf("    Schüler: %s %s\n", student.FirstName, student.LastName)
			}
		}
	}
}


