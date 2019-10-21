package  main 
	
import "fmt"

func main () {

	var poin = 8840.0

	if persen := poin/100; persen >= 100 {
		fmt.Printf("%.1f%s Sempurna\n", persen, "%")
		
	} else if persen >= 70 {
		fmt.Printf("%.1f%s Baik \n", persen, "%")
		
	} else {
		fmt.Printf("%.1f%s Lumayan \n", persen, "%")
	}
	
}