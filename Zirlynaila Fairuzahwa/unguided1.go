package main
import "fmt"

func main(){
	var berat, kg, gram, biaya, biayaTambahan, totalBiaya int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&berat)
	kg = berat / 1000
    gram = berat % 1000
	biaya = kg * 10000
	if gram < 1000 && kg > 10 {
		biayaTambahan = 0
	} else {
		if gram >= 500 {
			biayaTambahan = gram * 5
		} else {
			biayaTambahan = gram * 15
		}
	}
	totalBiaya = biaya + biayaTambahan
	fmt.Println("Detail berat: ", kg, " kg + ", gram, " gr")
	fmt.Println("Detail biaya: Rp. ", biaya, " + Rp. ", biayaTambahan)
	fmt.Println("Total biaya: Rp. ", totalBiaya)
}