package main

import "fmt"

func main() {
	var berat int
	fmt.Scan(&berat)

	kg := berat / 1000
	sisakg := berat % 1000
	biaya := 10000
	totalBiaya := 0

	if kg > 10 {
		totalBiaya = kg * biaya
	} else if sisakg >= 500 {
		totalBiaya = (biaya * kg) + (sisakg * 5)
	} else {
		totalBiaya = (biaya * kg) + (sisakg * 15)
	}

	fmt.Println(totalBiaya)
}
