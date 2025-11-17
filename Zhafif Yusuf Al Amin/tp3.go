package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	hasil1 := x / 1000
	hasil2 := x % 1000 / 100
	hasil3 := x % 100 / 10
	hasil4 := x % 10

	if hasil1 < hasil2 && hasil2 < hasil3 && hasil3 < hasil4 {
		fmt.Print("hasil terurut membesar")
	} else if hasil1 > hasil2 && hasil2 > hasil3 && hasil3 > hasil4 {
		fmt.Print("hasil terurut mengecil")
	} else {
		fmt.Print("tidak terurut")
	}

}
