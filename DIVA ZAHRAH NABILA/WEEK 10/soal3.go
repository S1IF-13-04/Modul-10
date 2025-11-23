package main

import "fmt"

func main() {
	var bil, faktor int
	fmt.Scan(&bil)
	faktor = 0

	for i := 1; i <= bil; i++ {
		if bil%i == 0 {
			fmt.Printf("%d ", i)
			faktor++
		}
	}
	fmt.Println()

	if faktor == 2 {
		fmt.Println("Prima:true")
	} else if faktor > 2 {
		fmt.Println("Prima:false")
	} else {
		fmt.Println("Prima:false")
	}
}
