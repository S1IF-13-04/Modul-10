package main

import "fmt"

func main() {
	var num int
	prime := 0
	fmt.Scan(&num)

	for i := num; i > 0; i-- {
		factor := num % i
		if factor == 0 {
			prime += 1
			fmt.Print(i, " ")
		}
	}
	if prime == 2 {
		fmt.Println("\n", num, "is a prime number")
	} else {
	}
}
