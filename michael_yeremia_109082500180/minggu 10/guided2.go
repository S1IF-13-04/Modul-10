package main
import "fmt"

func main() {
    var huruf rune
    fmt.Scanf("%c", &huruf)

    if huruf >= 65 && huruf <= 90 {
        huruf += 32
    }
    if huruf == 97 || huruf == 105 || huruf == 117 || huruf == 101 || huruf == 111 {
        fmt.Println("vokal")

    } else if huruf >= 97 && huruf <= 122 {
        fmt.Println("konsonan")

    } else {
        fmt.Println("bukan huruf")
    }
}