package main

import "fmt"

func main() {
	var x int
	var teks string
	fmt.Print("masukan :")
	fmt.Scan(&x)

	teks = "bukan"

	if x < 0 {
		x = x / 2
		teks = "genap negatif"
	}
	fmt.Print(teks)
}