package main

import (
    "fmt"
    "math/rand"
    "os"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    dice := rand.Intn(6) + 1
    timestamp := time.Now().Format(time.RFC1123)

    // TODO: Verwenden Sie os.Stdout und os.Stderr
    fmt.Fprintln(os.Stdout, dice)      // Würfelergebnis in Standardausgabe
    fmt.Fprintln(os.Stderr, timestamp) // Zeitpunkt in Standardfehler

    // TODO: Ergänzen Sie die Kommentare unten
    // Programmaufruf:
    // go run main.go > eyes.txt 2> dice.log
}
