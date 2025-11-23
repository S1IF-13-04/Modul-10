package main

import "fmt"

func main() {
    var nilai float64
    var huruf string

    fmt.Print("Nilai akhir mata kuliah: ")
    fmt.Scan(&nilai)

    if nilai >= 85 {
        huruf = "A"
    } else if nilai >= 70 {
        huruf = "B"
    } else if nilai >= 55 {
        huruf = "C"
    } else if nilai >= 40 {
        huruf = "D"
    } else {
        huruf = "E"
    }

    fmt.Println("Nilai mata kuliah:", huruf)
}
