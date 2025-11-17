package main

import "fmt"

func main() {
	var age int
	var famcard bool
	fmt.Scan(&age, &famcard)
	if age >= 17 && famcard {
		fmt.Println("can make an ID card")
	} else {
		fmt.Println("unable to make an ID card")
	}
}
