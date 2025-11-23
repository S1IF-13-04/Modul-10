package main

import "fmt"

func main() {
    var n float64
    var m string

    fmt.Scan(&n)

    if n > 80 {
        m = "A"
    } else if n > 72.5 {
        m = "AB"
    } else if n > 65 {
        m = "B"
    } else if n > 57.5 {
        m = "BC"
    } else if n > 50 {
        m = "C"
    } else if n > 40 {
        m = "D"
    } else {
        m = "E"
    }

    fmt.Println(m)
}
