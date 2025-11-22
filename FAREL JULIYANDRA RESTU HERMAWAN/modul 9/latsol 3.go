package main

import "fmt"

func main (){
	var x, y, a,b int
	fmt.Print("Masukan :")
	fmt.Scan(&x,&y)

	a = y % x
	b = x % y

	if a == 0{
		fmt.Println("true")
	}else {
		fmt.Println("false")
	}

	if b == 0 {
		fmt.Println("true")
	}else {
		fmt.Println("false")
	}
}