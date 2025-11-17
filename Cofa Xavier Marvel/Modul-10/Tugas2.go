package main

import "fmt"

func main() {
	var nam float64
	var nmk string
	fmt.Print("Final grade of the course: ")
	fmt.Scan(&nam)
	if nam > 80 {
		nam = "A"
	}
	if nam > 72.5 {
		nam = "AB"
	}
	if nam > 65 {
		nam = "B"
	}
	if nam > 57.5 {
		nam = "BC"
	}
	if nam > 50 {
		nam = "C"
	}
	if nam > 40 {
		nam = "D"
	} else if nam <= 40 {
		nam = "E"
	}
	fmt.Println("Course grades :", nmk)
}

//a.if NAM = 80.1	 the out put is Error
//  the nam is defined as float64 and not string
//  so when we assign nam = "A" it will cause an error

//b.the error of the program is that the variable nam is defined as float64
//  ,in the if statements it trys to assign a string value to the nam
//  variable to and the if statments are are ran in order and not in one 
//  if-else statment.
//  to fix the error we need use the variable nmk in the if statments
//  assigning the string values to nmk
//  and rewrite the if statments to if-else if statment

//c.

package main

import "fmt"

func main() {
	var nam float64
	var nmk string
	fmt.Print("Final grade of the course: ")
	fmt.Scan(&nam)
	if nam > 80 {
		nmk = "A"
	} else if nam > 72.5 {
		nmk = "AB"
	} else if nam > 65 {
		nmk = "B"
	} else if nam > 57.5 {
		nmk = "BC"
	} else if nam > 50 {
		nmk = "C"
	} else if nam > 40 {
		nmk = "D"
	} else if nam <= 40 {
		nmk = "E"
	}
	fmt.Println("Course grades :", nmk)
}