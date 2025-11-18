package main

import "fmt"

func main() {
	var bilangan int
	var status string

	fmt.Scan(&bilangan)

	d1 := bilangan / 1000
	d2 := (bilangan / 100) % 10
	d3 := (bilangan / 10) % 10
	d4 := bilangan % 10

	if d1 < d2 && d2 < d3 && d3 < d4 {
		status = "terurut membesar"
	} else if d1 > d2 && d2 > d3 && d3 > d4 {
		status = "terurut mengecil"
	} else {
		status = "tidak terurut"
	}

	fmt.Printf("Digit pada bilangan %d %s\n", bilangan, status)
}