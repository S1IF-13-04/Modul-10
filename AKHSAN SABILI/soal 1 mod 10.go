package main
import "fmt"
func main(){
	var gr, tambah int
	fmt.Print("berat parsel (gram) : ")
	fmt.Scan(&gr)
	kg := gr / 1000
	gram := gr % 1000
	fmt.Println("detail berat :", kg, "kg +", gram, "gr")
	biaya := kg * 10000
	if gram >= 500{
		tambah = gram * 5
	} else if gram < 500{
		tambah = gram * 15
	} 
	fmt.Println("detail biaya : Rp.", biaya, "+", tambah)
	if kg >= 10 {
		tambah = gram * 0
	}
	total := biaya + tambah
	 
	fmt.Println("total biaya :", total )
}