package main

import "fmt"

func main() {
	package main

import "fmt"

func outputDateRange(zodiacSign rune) {
    switch zodiacSign {
    case '♈': // Widder (Aries)
        fmt.Println("21.03. - 20.04.")
    case '♉': // Stier (Taurus)
        fmt.Println("21.04. - 20.05.")
    case '♊': // Zwillinge (Gemini)
        fmt.Println("21.05. - 21.06.")
    case '♋': // Krebs (Cancer)
        fmt.Println("22.06. - 22.07.")
    case '♌': // Löwe (Leo)
        fmt.Println("23.07. - 23.08.")
    case '♍': // Jungfrau (Virgo)
        fmt.Println("24.08. - 23.09.")
    case '♎': // Waage (Libra)
        fmt.Println("24.09. - 23.10.")
    case '♏': // Skorpion (Scorpio)
        fmt.Println("24.10. - 22.11.")
    case '♐': // Schütze (Sagittarius)
        fmt.Println("23.11. - 21.12.")
    case '♑': // Steinbock (Capricorn)
        fmt.Println("22.12. - 20.01.")
    case '♒': // Wassermann (Aquarius)
        fmt.Println("21.01. - 19.02.")
    case '♓': // Fische (Pisces)
        fmt.Println("20.02. - 20.03.")
    default:
        fmt.Println("Unknown zodiac sign.")
    }
}

func main() {
    // Beispiele für Sternzeichen
    outputDateRange('♈') // Widder
    outputDateRange('♋') // Krebs
    outputDateRange('♏') // Skorpion
    outputDateRange('♓') // Fische
    outputDateRange('❓') // Test für ein unbekanntes Zeichen
}

}