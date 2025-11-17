package main

import "fmt"

func main() {
	var x rune
	fmt.Scanf("%c", &x)

	letter := (x >= 'a' && x <= 'z') || (x >= 'A' && x <= 'Z')
	Vowel := (x == 'a' || x == 'i' || x == 'u' || x == 'e' || x == 'o' ||
		x == 'A' || x == 'I' || x == 'U' || x == 'E' || x == 'O')

	if letter && Vowel {
		fmt.Println("vowel!")
	} else if letter && !Vowel {
		fmt.Println("consonant!")
	} else {
		fmt.Println("non-letter!")
	}
}
