package main

import "fmt"

func main() {
	var Usia int
	var KK bool

	fmt.Scan(&Usia, &KK)

	if Usia >= 17 && KK == true {
		fmt.Print("Bisa membuat KTP")
	} else {
		fmt.Print("Tidak bisa membuat KTP")
	}
}
