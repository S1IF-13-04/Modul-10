package main

import "fmt"

func main() {
	var berat, kg, sisa, hargaTambahan, hargaAkhir int
	fmt.Scan(&berat)

	kg = berat / 1000
	sisa = berat % 1000
	fmt.Println("Detail berat:", kg, "kg +", sisa, "gr")
	hargaAkhir = kg * 10000

	if sisa >= 500 && kg <= 10 {
		hargaTambahan = sisa * 5
		fmt.Println("Detail biaya: Rp.", hargaAkhir, "+ Rp.", hargaTambahan)
		hargaAkhir = hargaAkhir + hargaTambahan
		fmt.Print("Total biaya: ", hargaAkhir)
	} else if sisa < 500 && kg <= 10 {
		hargaTambahan = sisa * 15
		fmt.Println("Detail biaya: Rp.", hargaAkhir, "+ Rp.", hargaTambahan)
		hargaAkhir = hargaAkhir + hargaTambahan
		fmt.Print("Total biaya: ", hargaAkhir)
	} else if kg > 10 {
		fmt.Println("Detail biaya: Rp.", hargaAkhir, "+ Rp.", hargaTambahan)
		fmt.Print("Total biaya: ", hargaAkhir)
	}
}
