package main 
import "fmt"

func main () {
	var karakter byte
	fmt.Scanf("%c",&karakter)

	if karakter >= 'A' && karakter <= 'Z'{
		karakter= karakter + 32
	}

	if karakter <97 || karakter> 122 {
		fmt.Println("bukan huruf")
	} else if karakter == 'a' || karakter == 'i' || karakter == 'u' || karakter == 'e' || karakter == 'o' {
		fmt.Println("vokal")
	} else {
		fmt.Println("konsonan")
	}
}