package main

import (
	"fmt"
)

func main() {
	var b,a,i int64
	var y,p bool

	fmt.Scan(&a)
	b = 0
	fmt.Print("Faktor: ")
	for i = 1 ; i <= a ; i ++ {
		y = a % i == 0
		if y == true{
			fmt.Printf(" %d ",i)
			b +=1
		}
	}
	p= b ==2

	fmt.Printf("\nPrima: %t",p)
}
