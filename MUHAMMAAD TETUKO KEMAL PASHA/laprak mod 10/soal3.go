package main

import "fmt"

func main() {
	var angka int
	var faktor []int
	fmt.Scan(&angka)
	for i := 1; i <= angka; i++ {
		if angka%i == 0 {
			fmt.Print(i, " ")
			faktor = append(faktor, i)
		}
	}
	if len(faktor) == 2 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
