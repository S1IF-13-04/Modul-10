package main

import "fmt"

func main() {
	var bil int
	var kk bool

	fmt.Scan(&bil, &kk)

	if bil >= 17 && kk == true {
		fmt.Print("bisa membuat ktp")
	} else {
		fmt.Print("tidak bisa membuat ktp")
	}
}
