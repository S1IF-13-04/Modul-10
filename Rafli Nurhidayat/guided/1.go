package main
import "fmt"
func main() {
	var umur int
	var kk bool

	fmt.Scan(&umur, &kk)
	if umur >= 17 && kk {
		fmt.Println("bisa membuat KTP")
	} else {
		fmt.Println("belum bisa membuat KTP")
	}
}