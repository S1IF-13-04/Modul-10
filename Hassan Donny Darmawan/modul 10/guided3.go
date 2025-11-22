package main
import "fmt"

func main (){
	var a int
	fmt.Scan(&a)
	b:= a/1000
	c:= (a%1000)/100
	d:= (a%100)/10
	e:= (a%10)

	if b>c && c>d && d>e {
		fmt.Print("mengecil")
	} else if b<c && c<d && d<e {
		fmt.Print("membesar")
	} else {
		fmt.Print("tidak mengurut")
	}
}