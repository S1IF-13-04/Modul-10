package main
import "fmt"
func main() {
	var b, i, jumlahFaktor int
	var isPrime bool
	fmt.Print("Bilangan: ")
	fmt.Scan(&b)
	fmt.Print("Faktor: ")
	for i = 1; i <= b; i++ {
		if b%i == 0 {
			fmt.Printf("%d ", i)
			jumlahFaktor++
		}
	}
	fmt.Println()
	if jumlahFaktor == 2 {
		isPrime = true
	} else {
		isPrime = false
	}
	fmt.Printf("Prima: %t\n", isPrime)
}