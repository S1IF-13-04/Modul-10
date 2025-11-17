package main

import "fmt"

func main() {
	var number, d1, d2, d3, d4 int
	var text string
	fmt.Scan(&number)

	d4 = number % 10
	d3 = number / 10 % 10
	d2 = number / 100 % 10
	d1 = number / 1000 % 10

	if d1 < d2 && d2 < d3 && d3 < d4 {
		text = "are in increasing order"
	} else if d1 > d2 && d2 > d3 && d3 > d4 {
		text = "are in decreasing order"
	} else {
		text = "are not sorted"
	}
	fmt.Println("The digits in number", number, text)
}
