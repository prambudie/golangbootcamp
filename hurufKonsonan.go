package main  
  
import "fmt"  
 
  
func tugasKonsonan(hurufVocal rune) {  
 if 
 hurufVocal == 'a' || 
 hurufVocal == 'e' ||
 hurufVocal == 'i' ||
 hurufVocal == 'o' || 
 hurufVocal == 'u' {  
  fmt.Printf(" %c adalah sebuah huruf vokal\n", hurufVocal)  
 } else {  
  fmt.Printf(" %c adalah sebuah huruf konsonan\n", hurufVocal)  
 }  
  
}  

func main() {  
	tugasKonsonan('g') // vokal 
	tugasKonsonan('b') // konsonan  
   } 
