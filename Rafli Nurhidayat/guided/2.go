package main

import "fmt"

func main() {
	var x string
	var huruf, hkecil, hkapital bool

	fmt.Scan(&x)

	huruf = (x >= "a" && x <= "z") || (x >= "A" && x <= "Z")
	hkecil = (x == "a" || x == "e" || x == "i" || x == "o" || x == "u")
	hkapital = (x == "A" || x == "E" || x == "I" || x == "O" || x == "U")

	if huruf && (hkecil || hkapital) {
		fmt.Println("vokal")
	} else if huruf && !(hkecil || hkapital) {
		fmt.Println("konsonan")
	} else {
		fmt.Println("bukan huruf")
	}


}
