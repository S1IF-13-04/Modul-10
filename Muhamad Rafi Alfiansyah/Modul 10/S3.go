package main

import "fmt"

func main() {
	var bil, i, f int
	var prima bool = true

	fmt.Print("Bilangan : ")
	fmt.Scan(&bil)
	fmt.Print("Faktor : ")

	if bil > 1 {
		for i = 1; i <= bil; i++ {
			if bil%i == 0 {
				fmt.Print(i, " ")
				f++
			}
		}
	}

	if f == 2 {
		fmt.Println("\nPrima : ", prima)
	} else {
		fmt.Println("\nPrima : ", !prima)
	}
}
