package main

import "fmt"

func main() {
	var berat, kg, gr int
	var biayaKg, biayaGr, totalbiaya int

	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&berat)

	kg = berat / 1000
	gr = berat % 1000
	biayaKg = kg * 10000

	if kg > 10 {
		biayaGr = 0
	} else if gr >= 500 {
		biayaGr = gr * 5
	} else {
		biayaGr = gr * 15
	}

	totalbiaya = biayaKg + biayaGr

	fmt.Printf("Detail berat: %d kg + %d gr\n", kg, gr)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKg, biayaGr)
	fmt.Printf("Total biaya: Rp. %d\n", totalbiaya)
}
