package main
import "fmt"
func main(){
	var b, t int
	var prima bool
	fmt.Scan(&b)
	for i := 1; i <= b; i++{
		if b % i == 0{
			fmt.Print(i," ")
			t++
		}
	} 
	if t == 2{
		prima = true
	} else {
		prima = false
	}
	fmt.Printf("\n Prima : %t", prima)
}