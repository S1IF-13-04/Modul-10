package main

import "fmt"

func main() {
	var huruf string

	fmt.Scan(&huruf)

	if huruf == "A" || huruf == "a" || huruf == "I" || huruf == "i" || huruf == "u" || huruf == "E" ||
		huruf == "e" || huruf == "O" || huruf == "o" {
		fmt.Print("Vokal")
	} else if (huruf >= "A" && huruf <= "Z") || (huruf >= "a" && huruf <= "z") {
		fmt.Print("Konsonan")
	} else {
		fmt.Print("bukan huruf")
	}

}
