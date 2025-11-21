package main

import "fmt"

func main() {
	var bil, d1, d2, d3, d4 int
	fmt.Scan(&bil)

	d1 = bil / 1000
	d2 = (bil % 1000) / 100
	d3 = (bil % 100) / 10
	d4 = (bil % 10)

	if d1 < d2 && d2 < d3 && d3 < d4 {
		fmt.Println("Terurut membesar")
	} else if d1 > d2 && d2 > d3 && d3 > d4 {
		fmt.Println("Terurut mengecil")
	} else {
		fmt.Println("Tidak terurut")
	}
}
