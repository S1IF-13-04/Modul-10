package main

import "fmt"

func main() {

	var angka int
	var teks string

	fmt.Scan(&angka)

	d1 := angka / 1000
	d2 := (angka / 100) % 10
	d3 := (angka / 10) % 10
	d4 := angka % 10

	if d1 > d2 && d2 > d3 && d3 > d4 {
		teks = "terurut mengecil"
	} else if d1 < d2 && d2 < d3 && d3 < d4 {
		teks = "terurut membesar"
	} else {
		teks = "tidak terurut"
	}
	fmt.Println("Digit pada bilangan", angka, teks)
}
	