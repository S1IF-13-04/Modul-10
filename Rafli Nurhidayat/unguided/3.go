package main

import "fmt"

func main() {
	var b, faktor int
	var prima bool

	fmt.Scan(&b)

	fmt.Println("Bilangan: ", b)
	fmt.Print("Faktor: ")

	for i := 1; i <= b; i++ {
		if b % i == 0 {
			fmt.Print(i, " ")
		faktor++
		}
	}

	if faktor == 2 {
		fmt.Println("\nPrima:", prima)
	} else {
		fmt.Println("\nPrima:", !prima)
	}
}