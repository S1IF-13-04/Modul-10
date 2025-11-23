package main
import "fmt"

func main() {
	var ktp int
	var kk bool

	fmt.Scan(&ktp, &kk)
	if ktp >= 17 && kk {
		fmt.Print("bisa membuat ktp")
	}else{
		fmt.Print("belum bisa membuat ktp")
	}
}