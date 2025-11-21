package main

import ("fmt" 
		"math")

func main() {
	var berat, biayaGram float64

	fmt.Scan(&berat)

	kilo := math.Floor(berat / 1000)
	gram := math.Mod(berat, 1000)

	if kilo > 10 {
		biayaGram = 0
	} else if gram >= 500 {
		biayaGram = gram * 5
	} else {
		biayaGram = gram * 15
	}

	biayaKilo := kilo * 10000
	total := biayaKilo + biayaGram
	
	fmt.Println("Berat Parsel (gram):", berat)
	fmt.Println("Detail Berat:", kilo, "kg +", gram, "gr")
	fmt.Println("Detail Biaya: Rp.", biayaKilo, "+ Rp", biayaGram)
	fmt.Println("Total Biaya: Rp.", total)
}