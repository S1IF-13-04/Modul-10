package main
import "fmt"

func main() {
    var angka,a,b,c,d int
    fmt.Scan(&angka)

	a = angka /1000
	b = angka /100 %10
	c = angka /10 %10
	d = angka %10

    if a>b && b>c && c>d{
        fmt.Printf("digit pada bilangan %d terurut mengecil", angka)
    }else if a<b && b<c && c<d{
        fmt.Printf("digit pada bilangan %d terurut membesar",angka)
	}else{
        fmt.Printf("digit pada bilangan %d tidak terurut ", angka)
	}
}