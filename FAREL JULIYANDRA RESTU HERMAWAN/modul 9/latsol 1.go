package main

import "fmt"

func main() {
	var n, m, s int
	fmt.Print("Masukan :")
	fmt.Scan(&n)

	m = n / 2
	s = n % 2

	if s != 0 {
		m++
	}
	fmt.Print("Keluaran : ",m)
}