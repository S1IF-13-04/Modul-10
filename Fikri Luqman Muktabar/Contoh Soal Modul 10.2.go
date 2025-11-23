package main  
import "fmt"  
func main() {      
	var x rune 
    var huruf, vKecil, vKapital bool     
	fmt.Scanf("%c", &x) 
    huruf = (x >= 'a' && x <= 'z') || (x >= 'A' && x <= 'Z')     
	vKecil = x == 'a' || x == 'i' || x == 'u' || x == 'e' || x == 'o'     
	vKapital = x == 'A' || x == 'I' || x == 'U' || x == 'E' || x == 'O'     
	if huruf && (vKecil || vKapital) {          
		fmt.Println("Vokal")  
    }else if huruf && !(vKecil || vKapital) {         
		fmt.Println("Konsonan") 
    }else{  
        fmt.Println("Bukan huruf")     
	} 
 }  