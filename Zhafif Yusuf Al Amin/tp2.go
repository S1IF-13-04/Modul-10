package main

import "fmt"

func main() {
	var x rune
	fmt.Scanf("%c", &x)
	if (x >= 'a' && x <= 'z') || (x >= 'A' && x <= 'Z') {
		if (x == 'a' || x == 'i' || x == 'u' || x == 'e' || x == 'o') ||
			(x == 'A' || x == 'I' || x == 'U' || x == 'E' || x == 'O') {
			fmt.Print("vokal")
		} else {
			fmt.Print("konsonan")
		}
	} else {
		fmt.Print("bukan huruf")
	}
}
