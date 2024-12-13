package main

import "fmt"

const (
	Aries       = "♈" 
	Taurus      = "♉"
	Gemini      = "♊"
	Cancer      = "♋" 
	Leo         = "♌"
	Virgo       = "♍" 
	Libra       = "♎" 
	Scorpio     = "♏" 
	Sagittarius = "♐" 
	Capricorn   = "♑" 
	Aquarius    = "♒" 
	Pisces      = "♓" 
)

type Person struct {
	Name  string
	Day   int
	Month int
}

func outputWithZodiacSign(p Person) {
	var zodiacSign string
	if (p.Month == 3 && p.Day >= 21) || (p.Month == 4 && p.Day <= 20) {
		zodiacSign = Aries
	} else if (p.Month == 4 && p.Day >= 21) || (p.Month == 5 && p.Day <= 20) {
		zodiacSign = Taurus
	} else if (p.Month == 5 && p.Day >= 21) || (p.Month == 6 && p.Day <= 21) {
		zodiacSign = Gemini
	} else if (p.Month == 6 && p.Day >= 22) || (p.Month == 7 && p.Day <= 22) {
		zodiacSign = Cancer
	} else if (p.Month == 7 && p.Day >= 23) || (p.Month == 8 && p.Day <= 23) {
		zodiacSign = Leo
	} else if (p.Month == 8 && p.Day >= 24) || (p.Month == 9 && p.Day <= 23) {
		zodiacSign = Virgo
	} else if (p.Month == 9 && p.Day >= 24) || (p.Month == 10 && p.Day <= 23) {
		zodiacSign = Libra
	} else if (p.Month == 10 && p.Day >= 24) || (p.Month == 11 && p.Day <= 22) {
		zodiacSign = Scorpio
	} else if (p.Month == 11 && p.Day >= 23) || (p.Month == 12 && p.Day <= 21) {
		zodiacSign = Sagittarius
	} else if (p.Month == 12 && p.Day >= 22) || (p.Month == 1 && p.Day <= 20) {
		zodiacSign = Capricorn
	} else if (p.Month == 1 && p.Day >= 21) || (p.Month == 2 && p.Day <= 19) {
		zodiacSign = Aquarius
	} else if (p.Month == 2 && p.Day >= 20) || (p.Month == 3 && p.Day <= 20) {
		zodiacSign = Pisces
	}

	fmt.Printf("%s, born on %02d.%02d, has the zodiac sign %s.\n", p.Name, p.Day, p.Month, zodiacSign)
}

func main() {
	people := []Person{
		{"Grace Hopper", 9, 12},
		{"Dennis Ritchie", 9, 9},
		{"Rick Astley", 6, 2},
		{"Edsger Dijkstra", 11, 5},
		{"Alan Turing", 23, 6},
	}

	for _, person := range people {
		outputWithZodiacSign(person)
	}
}
