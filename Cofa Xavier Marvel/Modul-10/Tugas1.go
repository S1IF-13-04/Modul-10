package main

import "fmt"

func main() {
	var grams int64
	fmt.Scan(&grams)

	kg := grams / 1000
	g := grams % 1000
	price := kg * 10000

	if g >= 500 && kg < 10 {
		price = price + (g * 5)
	} else if g < 500 && kg < 10 {
		price = price + (g * 15)
	}

	fmt.Println(price)
}
