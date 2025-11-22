package main

import "fmt"

func main() {
	var h rune
	fmt.Scanf("%c", &h)

	if h == 'A' || h == 'I' || h == 'U' || h == 'E' || h == 'O' ||
		h == 'a' || h == 'i' || h == 'u' || h == 'e' || h == 'o' {
		fmt.Println("vokal")
	} else if (h >= 'a' && h <= 'z') || (h > 'A' && h <= 'Z') {
		fmt.Println("konsonan")
	} else {
		fmt.Println("bukan huruf")
	}

}
