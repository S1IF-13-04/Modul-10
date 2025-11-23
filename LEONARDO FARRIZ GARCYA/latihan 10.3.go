package main

import "fmt"

func main() {
    var b int
    fmt.Scan(&b)

    fmt.Print("Faktor: ")
    prima := true

    for i := 1; i <= b; i++ {
        if b%i == 0 {
            fmt.Printf("%d ", i)
            if i != 1 && i != b {
                prima = false
            }
        }
    }
    fmt.Printf("\nPrima: %v\n", prima)
}
