package main
import "fmt"
func main() {
    var b int
    var i int
    var count int
    var prima bool

    fmt.Print("bilangan: ")
    fmt.Scan(&b)

    fmt.Print("faktor: ")
    count = 0
    for i = 1; i <= b; i++ {
        if b % i == 0 {
            fmt.Print(i, " ")
            count = count + 1
        }
    }
    fmt.Println()

    prima = false
    if count == 2 {
        prima = true
    }

    fmt.Println("prima:", prima)
}
