package main

import "fmt"

func main() {
	var bil, i, sisa, faktor int
	var prima bool
	fmt.Scan(&bil)

	for i = 1; i <= bil; i++ {
		sisa = bil % i
		if sisa == 0 {
			faktor = i
			fmt.Print(faktor, " ")
		}
	}
	fmt.Println()

	if bil%2 == 0 || bil%3 == 0 {
		prima = false
	} else {
		prima = true
	}
	if bil <= 3 && bil > 0 {
		prima = true
	}
	fmt.Print("Prima: ", prima)
}
