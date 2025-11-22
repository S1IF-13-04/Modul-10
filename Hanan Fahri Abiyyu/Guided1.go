package main

import "fmt"

func main() {
	var umur int
	var KK bool
	fmt.Scan(&umur, &KK)

	if umur >= 17 && KK == true{
		fmt.Print("bisa membuat ktp")
	}else{
		fmt.Print("belum bisa")
	}
}
