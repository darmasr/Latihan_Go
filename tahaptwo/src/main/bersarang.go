package  main 
	
import "fmt"

func main () {

	var poin = 0

	if poin > 7  {
		switch poin{
		case 10:
			fmt.Println("Sempurna")
		default :
			fmt.Println("Bagus")
		}
		
	} else {
		if poin == 5 {
			fmt.Println("Lumayan")
			
		} else if poin == 3 {
			fmt.Println("Tetap Mencoba")
			
		} else {
			fmt.Println("Kamu pasti bisa")
			if poin == 0 {
				fmt.Println("Coba lebih giat lagi")
				
			}
		}
	}

	
}