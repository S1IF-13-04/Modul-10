package main

import "fmt"

func main() {
	var x int
	var y int = 0
	fmt.Print("Bilangan: ")
	fmt.Scan(&x)

	fmt.Print("Faktor: ")

	for i := 1; i <= x; i++ {
		if x%i == 0 {
			fmt.Print(i, " ")
			y++
		}
	}
	fmt.Println()

	if y == 2 {
		fmt.Println("Prima: true")
	} else {
		fmt.Println("Prima: false")
	}
}
