package main

import "fmt"

func main() {
	var bilangan int
	var d1, d2, d3, d4 int
	fmt.Scan(&bilangan)
	d1 = bilangan / 1000
	d2 = bilangan / 100 % 10
	d3 = bilangan / 10 % 10
	d4 = bilangan % 10

	if d1 < d2 && d2 < d3 && d3 < d4 {
		fmt.Print("Digit pada bilangan ", bilangan, " terurut membesar")
	} else if d1 > d2 && d2 > d3 && d3 > d4 {
		fmt.Print("Digit pada bilangan ", bilangan, " terurut mengecil")
	} else {
		fmt.Print("Digit pada bilangan ", bilangan, " tidak terurut")
	}
}
