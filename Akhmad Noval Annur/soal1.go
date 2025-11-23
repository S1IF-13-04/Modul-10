package main

import "fmt"

func main() {
    var g, k, s int
    var bkg, bs, t int

    fmt.Print("Berat parsel (gram): ")
    fmt.Scan(&g)

    k = g / 1000
    s = g % 1000

    bkg = k * 10000

    if k > 10 {
        bs = 0
    } else {
        if s >= 500 {
            bs = s * 5
        } else {
            bs = s * 15
        }
    }

    t = bkg + bs

    fmt.Printf("Detail berat: %d kg + %d gr\n", k, s)
    fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", bkg, bs)
    fmt.Printf("Total biaya: Rp. %d\n", t)
}
