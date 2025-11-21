package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Print("Bilangan: ")
	fmt.Scan(&a)

	fmt.Print("Faktor: ")

	var b int = 0

	for i := 1; i <= a; i++ {
		if a%i == 0 {
			fmt.Print(i, " ")
			b++
		}
	}
	fmt.Println()

	if b == 2 {
		fmt.Println("Prima: true")
	} else {
		fmt.Println("Prima: false")
	}
}
