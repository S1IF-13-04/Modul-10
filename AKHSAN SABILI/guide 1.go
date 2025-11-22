package main
import "fmt"
func main(){
	var a int
	var kk bool
	fmt.Scan(&a, &kk)
	if a >= 17 && kk == true{
		fmt.Println("bisa membuat ktp")
	} else {
		fmt.Println("tidak bisa membuat ktp")
	}
}