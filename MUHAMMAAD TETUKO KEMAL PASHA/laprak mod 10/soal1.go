package main

import "fmt"

func main() {
	var parsel int
	var kgParsel int
	var total int
	var sisa int
	fmt.Scan(&parsel)
	kgParsel = parsel / 1000
	sisa = parsel % 1000
	if kgParsel > 0 {
		total += kgParsel * 10000
		if kgParsel < 10 {
			if sisa >= 500 {
				total += sisa * 5
			} else if sisa < 500 {
				total += sisa * 15
			}
		}
	}
	fmt.Print(total)
}
