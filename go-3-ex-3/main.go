package main

import "fmt"

const (
    Lower = 1  // Startwert
    Upper = 100 // Endwert
)

func main() {
    for i := Lower; i <= Upper; i++ {
        switch {
        case i%3 == 0 && i%5 == 0: // Durch 3 und 5 teilbar
            fmt.Println("FizzBuzz")
        case i%3 == 0: // Durch 3 teilbar
            fmt.Println("Fizz")
        case i%5 == 0: // Durch 5 teilbar
            fmt.Println("Buzz")
        default: // Für alle anderen Zahlen
            fmt.Println(i)
        }
    }
}
