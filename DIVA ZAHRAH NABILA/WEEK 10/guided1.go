package main

import "fmt"

func main() {

	var usia int
	var KK bool

	fmt.Scan(&usia, &KK)

	if usia >= 17 && KK {
		fmt.Print("bisa membuat KTP")
	} else {
		fmt.Print("belum bisa membuat KTP")
	}
}
