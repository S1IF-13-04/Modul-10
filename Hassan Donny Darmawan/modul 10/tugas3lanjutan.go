package main
import "fmt"

func main (){
	var b,i,jumlah_f int
	fmt.Scan(&b)
	jumlah_f = 0

	for i=1; i<=b;i++ {
		if b%i==0 {
			fmt.Print(i," ")
			jumlah_f++
		}
	}
	fmt.Println()
	
	if jumlah_f==2 {
		fmt.Print("merupakan bilangan Prima")
	}else{
		fmt.Print("bukan bilangan prima")
	}
}