package main

import "fmt"

func main() {
	var gram int
	fmt.Scan(&gram)

	kg := gram / 1000
	sisa := gram % 1000
	biaya := kg * 10000
	tambahan := 0
	if kg > 10 {
		tambahan = 0
	} else if sisa >= 500 {
		tambahan = sisa * 5
	} else if sisa > 0 {
		tambahan = sisa * 15
	} else {
		tambahan = 0
	}
	total := biaya + tambahan
	fmt.Println("Berat parsel (gram):", gram)
	fmt.Println("Detail berat:", kg, "+", sisa)
	fmt.Println("Detail biaya:", biaya, "+", tambahan)
	fmt.Println("Total biaya pengiriman: ", total)
}
