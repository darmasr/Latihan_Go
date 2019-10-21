package  main 
	
import "fmt"

func main () {

	var poin = 4
	if poin == 10 {
		fmt.Println("lulus dengan nilai sempurna")
		
	} else if poin >5 {
		fmt.Println("lulus")
		
	} else if poin == 4 {
		fmt.Println("hampir lulus")
		
	} else{
		fmt.Printf("tidak lulus. nilai anda %d\n", poin)
	}
	
}