package main

import "fmt"

func main() {
    suits := []rune{'♠', '◆', '♣', '♥'} // Farben: Karo, Pik, Kreuz, Herz
    ranks := []string{"⑥", "⑦", "⑧", "⑨", "⑩", "J", "Q", "K", "A"} // Werte: 6 bis Ass

    for _, rank := range ranks { // Äußere Schleife: Werte (ranks)
        for _, suit := range suits { // Innere Schleife: Farben (suits)
            fmt.Printf("%c%s\t", suit, rank) // Karte ausgeben (Symbol + Wert)
        }
         fmt.Println() // Neue Zeile nach jeder Reihe von Farben.
    }
}
