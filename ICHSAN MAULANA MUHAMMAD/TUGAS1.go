package main

import "fmt"

func main() {
	var gram, sisa, kg, total int
	fmt.Print("Berat Parsel: ")
	fmt.Scan(&gram)
	kg = gram / 1000
	sisa = gram % 1000
	fmt.Println("Detail Berat:", kg, "kg", "+", sisa, "gr")
	kg = kg * 10000
	if kg > 100000 {
		sisa = sisa * 0
	} else if kg <= 100000 && sisa >= 500 {
		sisa = sisa * 5
	} else if kg <= 100000 && sisa < 500 {
		sisa = sisa * 15
	}
	fmt.Println("Detail Biaya:", "Rp.", kg, "+", "Rp.", sisa)
	total = kg + sisa
	fmt.Println("Total Biaya:", "Rp.", total)
}
