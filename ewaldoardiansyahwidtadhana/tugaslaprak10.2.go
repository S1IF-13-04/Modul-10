package main

import "fmt"

func main() {
	var nam float64
	var nnk string

	fmt.Print("Nilai akhir mata kuliah: ")
	fmt.Scan(&nam)

	if nam >= 80 {
		nnk = "A"
	} else if nam > 72.5 && nam < 80 {
		nnk = "AB"
	} else if nam > 65 && nam <= 72.5 {
		nnk = "B"
	} else if nam > 57.5 && nam <= 65 {
		nnk = "BC"
	} else if nam > 50 && nam <= 57.5 {
		nnk = "C"
	} else if nam > 40 && nam <= 50 {
		nnk = "D"
	} else {
		nnk = "E"
	}

	fmt.Println("Nilai mata kuliah: ", nnk)
}
