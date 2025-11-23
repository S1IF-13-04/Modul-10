package main
import "fmt"
func main (){
	var b, i int
	fmt.Scan(&b)

	for i=i; i<=b;i++ {
        if b%i==0 {
            fmt.Print(i," ")
        }
    }
}