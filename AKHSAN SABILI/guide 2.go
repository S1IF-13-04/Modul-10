package main
import "fmt"
func main(){
	var n rune
	var huruf, hk, hb bool
	fmt.Scanf("%c", &n)
	huruf = (n >= 'a' && n <= 'z') || (n >= 'A' && n <= 'Z')
	hk = n == 'a' || n == 'i' || n == 'u' || n == 'e' || n == 'o'
	hb = n == 'A' || n == 'I' || n == 'U' || n == 'E' || n == 'O'
	if huruf && (hk || hb) {
		fmt.Println("vokal")
	}else if huruf && !(hk || hb){
		fmt.Println("konsonan")
	}else{
		fmt.Println("bukan huruf")
	}
}

