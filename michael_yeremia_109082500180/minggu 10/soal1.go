package main

import (
	"fmt"
)

func main() {
	var berat, kg, gr, harga, biaya, total int64
	fmt.Scan(&berat)
	kg = berat / 1000
	gr = berat % 1000

	harga = kg * 10000

	if kg >= 10 {
		biaya = 0
	} else {
		if gr < 500 {
			biaya = gr * 15
		} else if gr >= 500 {
			biaya = gr * 5
		}
	}
	total = harga + biaya

	fmt.Printf("Total biaya: Rp.%d", total)

}
