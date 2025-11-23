package main

import "fmt"

func main() {
	var kons string
	fmt.Scan(&kons)

	if kons == "A" || kons == "I" || kons == "U" || kons == "E" || kons == "O" ||
		kons == "a" || kons == "i" || kons == "u" || kons == "e" || kons == "o" {
		fmt.Print("Vokal")
	} else if (kons >= "a" && kons <= "z") || (kons >= "A" && kons <= "Z") {
		fmt.Print("Konsonan")
	} else {
		fmt.Print("Bukan huruf")
	}
}
