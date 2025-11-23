package main

import "fmt"

func main() {
	var bilangan int
	fmt.Print("Bilangan: ")
	fmt.Scan(&bilangan)

	fmt.Print("Faktor: ")
	faktor := 0

	for i := 1; i<= bilangan; i++ {
		if bilangan%i == 0 {
			fmt.Print(i, " ")
			faktor++
		}
	}
	fmt.Println()
	if faktor == 2 {
		fmt.Print("Prima: True")
	} else {
		fmt.Print("Prima: False")
	}
}