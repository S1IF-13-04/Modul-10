package main
import "fmt"
func main() {
	var beratTotal, kg, sisaGram, biayaKg, biayaSisa, totalBiaya int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&beratTotal)
	kg = beratTotal / 1000
	sisaGram = beratTotal % 1000
	biayaKg = kg * 10000
	if sisaGram >= 500 {
		biayaSisa = sisaGram * 5
	} else {
		biayaSisa = sisaGram * 15
	}
	if beratTotal > 10000 {
		biayaSisa = 0
	}
	totalBiaya = biayaKg + biayaSisa
	fmt.Printf("Detail berat: %d kg + %d gr\n", kg, sisaGram)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKg, biayaSisa)
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}