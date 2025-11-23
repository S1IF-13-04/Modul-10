package main

import "fmt"

func main() {
    var b int
    fmt.Scan(&b)

    c := 0

    for i := 1; i <= b; i++ {
        if b%i == 0 {
            fmt.Print(i, " ")
            c++
        }
    }

    fmt.Println()

    if c == 2 {
        fmt.Println("true")
    } else {
        fmt.Println("false")
    }
}
