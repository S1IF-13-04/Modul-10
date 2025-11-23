package main

import "fmt"

func main(){
var parsel int
    var biaya_kg, biaya_sisa, total float64
    fmt.Print("Berat parsel (gram): ")
    fmt.Scan(&parsel)

    berat := parsel/1000
    sisa := parsel % 1000
    biaya_kg = (float64(berat)*10000)

    if berat <10 && sisa <500 {
        biaya_sisa= (float64(sisa)*15) 
    } else if berat <10 && sisa>=500 {
        biaya_sisa= (float64(sisa)*5) 
    } else if berat>=10 && sisa <1000 {
        biaya_sisa= (float64(sisa)*0)
    }
    total= biaya_kg + biaya_sisa

    fmt.Println("Detail Berat:", berat, "kg +",sisa, "gr")
    fmt.Println("Detail biaya: Rp.", biaya_kg,"+ Rp.",biaya_sisa)
    fmt.Println("Total biaya: RP.", total)
}
