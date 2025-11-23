package main

import "fmt"

func main() {
    var berat int
    fmt.Print("Berat parsel (gram): ")
    fmt.Scan(&berat)

    kg := berat / 1000
    sisa := berat % 1000
    total := kg * 10000

    if kg < 10 {
        if sisa >= 500 {
            total += sisa * 5
        } else {
            total += sisa * 15
        }
    }

    fmt.Printf("Detail berat: %d kg %d gr\n", kg, sisa)
    fmt.Printf("Total biaya: Rp. %d\n", total)
}
