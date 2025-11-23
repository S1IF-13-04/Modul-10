package main

import "fmt"

func main() {
	var b int
	fmt.Print("Bilangan: ")
	fmt.Scan(&b)

	fmt.Print("Faktor: ")
	for f := 1; f <= b; f++ {
		if b%f == 0 {
			fmt.Print(f, " ")
		}
	}
}
