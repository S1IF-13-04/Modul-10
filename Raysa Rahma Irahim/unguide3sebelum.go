package main

import "fmt"

func main() {

	var b, f int

	fmt.Print("Bilangan: ")
	fmt.Scan(&b)

	fmt.Print("Faktor: ")

	for f = 1; f <= b; f++ {
		if b%f == 0 {
			fmt.Print(f, " ")
		}
	}
}
