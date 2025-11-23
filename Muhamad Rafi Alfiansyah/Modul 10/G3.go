package main
import "fmt"
func main(){
	var bil int
	fmt.Scan(&bil)
	a := bil/1000
	b := bil/100%10
	c := bil/10%10
	d := bil%10

	 if a < b && b < c && c < d {
        fmt.Println("digit terurut membesar")
    } else if a > b && b > c && c > d {
        fmt.Println("digit terurut mengecil")
    } else {
        fmt.Println("digit tidak terurut")
    }
}