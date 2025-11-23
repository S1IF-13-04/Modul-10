package main

import "fmt"

func main() {
	var ba, k, s, tk, ts, ta int64
	fmt.Print("Berat parsel (gram) : ")
	fmt.Scan(&ba)

	k = ba / 1000
	tk = k * 10000
	s = ba - (k * 1000)

	if s >= 500 {
		ts = s * 5
	} else {
		ts = s * 15
	}

	if k > 10 {
		ta = tk
	} else {
		ta = tk + ts
	}

	fmt.Printf("Detail berat : %d kg + %d gr \n", k, s)
	fmt.Printf("Detail biaya : Rp. %d + Rp. %d \n", tk, ts)
	fmt.Print("Total biaya : Rp. ", ta)
}
