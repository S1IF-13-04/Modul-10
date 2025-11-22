package main
import "fmt"

func main (){
	var nam float64
	var grade string
	fmt.Print("masukan nilai akhir mata kuliah: ")
	fmt.Scan(&nam)

	if  nam>80 {
		grade= "A"
	} else if nam>72.5{
		grade="AB"
	} else if nam>65{
		grade= "B"
	} else if nam>57.5{
		grade= "BC"
	} else if nam>50{
		grade="C"
	} else if nam>40{
		grade= "D"
	} else if nam<=40 {
		grade="E"
	}
	fmt.Println("Grade Nilai:",grade)
}