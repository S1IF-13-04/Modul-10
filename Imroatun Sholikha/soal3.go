package main

import "fmt"

func main() {
	var b, jumlahFaktor int
	jumlahFaktor = 0

	fmt.Print("Bilangan: ")
	fmt.Scan(&b)

	fmt.Print("Faktor: ")

	for i := 1; i <= b; i++ {
		if b%i == 0 {
			fmt.Print(i)
			if i < b {
				fmt.Print(" ")
			}
			jumlahFaktor = jumlahFaktor + 1
		}
	}

	fmt.Println()
	Prima := false

	if jumlahFaktor == 2 {
		Prima = true
	} else {
		Prima = false
	}

	if Prima {
		fmt.Println("Prima: true")
	} else {
		fmt.Println("Prima: false")
	}
}
