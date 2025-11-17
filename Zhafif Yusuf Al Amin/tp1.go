package main

import "fmt"

func main() {
	var usia int
	var kk bool
	fmt.Scan(&usia, &kk)

	if usia >= 17 && kk == true {
		fmt.Print("bisa membuat KTP")
	} else {
		fmt.Print("belum bisa membuat KTP")
	}
}
