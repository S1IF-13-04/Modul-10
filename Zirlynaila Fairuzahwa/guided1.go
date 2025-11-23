package main
import "fmt"

func main(){
	var usia int
	var kartuKeluarga bool
	fmt.Scan(&usia, &kartuKeluarga)
	if usia >= 17 && kartuKeluarga == true {
		fmt.Print("bisa membuat KTP")
	}else{
		fmt.Print("belum bisa membuat KTP")
	}
}