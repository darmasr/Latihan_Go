package  main 
	
import "fmt"

func main () {

	var poin = 6

	switch  {
	case poin == 8:
		fmt.Println("Sempurna")
	case (poin <8) && (poin > 5):
		fmt.Println("Baik")
		fallthrough
	case poin < 5:
		fmt.Println("Perbanyak belajar ya")
	default :
		fmt.Println("Lumayan")
		fmt.Println("Perbanyak belajar ya")
	}
	
}