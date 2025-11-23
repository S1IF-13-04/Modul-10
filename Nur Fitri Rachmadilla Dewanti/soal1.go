package main
import "fmt"
func main() {
    var berat, kg, sisa, biayaKg, biayaSisa, total int
    fmt.Print("berat parsel (gram): ")
    fmt.Scan(&berat)

    kg = berat / 1000
    sisa = berat % 1000
    fmt.Println("detail berat:", kg, "kg +", sisa, "gr")

    biayaKg = kg * 10000
    biayaSisa = 0

    if berat > 10000 {
        biayaSisa = 0
    } else {
        if sisa >= 500 {
            biayaSisa = sisa * 5
        } else {
            biayaSisa = sisa * 15
        }
    }

    total = biayaKg + biayaSisa
    fmt.Println("detail biaya: Rp.", biayaKg, "+ Rp.", biayaSisa)
    fmt.Println("total biaya: Rp.", total)
}
