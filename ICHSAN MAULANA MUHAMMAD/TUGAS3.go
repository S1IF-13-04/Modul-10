package main

import "fmt"

func main() {
	var b int
	fmt.Print("Bilangan: ")
	fmt.Scan(&b)

	fmt.Print("Faktor: ")
	a := 0
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			fmt.Print(i, " ")
			a++
		}
	}
	fmt.Println()

	if a == 2 {
		fmt.Println("Prima: true")
	} else {
		fmt.Println("Prima: false")
	}
}
