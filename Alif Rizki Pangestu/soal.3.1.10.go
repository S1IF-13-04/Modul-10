package main
import "fmt"

func main (){
    var b,i int
    fmt.Scan(&b)

    for i=1; i<=b;i++ {
        if b%i==0 {
            fmt.Print(i," ")
        }
    }
}
