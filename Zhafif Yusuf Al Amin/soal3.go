package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	jumlahFaktor := 0

	for i := 1; i <= x; i++ {
		if x%i == 0 {
			fmt.Print(i, " ")
			jumlahFaktor++
		}
	}
	fmt.Println()
	if jumlahFaktor == 2 {
		fmt.Print("PRIMA : TRUE")
	} else {
		fmt.Print("PRIMA : FALSE")
	}
}
