package main

import "fmt"

func main() {
	var beratParsel, totalBiaya, biayaPerGram, biayaTambahan int
	fmt.Print("Berat Parsel: ")
	fmt.Scan(&beratParsel)

	beratKg := beratParsel / 1000
	sisaGram := beratParsel % 1000
	biayaBerat := beratKg * 10000

	gratisSisa := beratParsel >= 10000

	if sisaGram > 0 && !gratisSisa {
		biayaPerGram = 0
		if sisaGram >= 500 {
			biayaPerGram = 5
			biayaTambahan = sisaGram * biayaPerGram
		} else {
			biayaPerGram = 15
			biayaTambahan = sisaGram * biayaPerGram
		}

	}
	fmt.Printf("Detail berat: %d kg + %d gr\n", beratKg, sisaGram)

	totalBiaya = biayaBerat + biayaTambahan

	fmt.Print("Detail biaya: ")
	if gratisSisa && sisaGram > 0 {
		fmt.Printf("Rp. %d + Rp. 0 (Gratis sisa %d gr)\n", biayaBerat, sisaGram)
	} else {
		fmt.Printf("Rp. %d + Rp. %d\n", biayaBerat, biayaTambahan)
	}

	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}
