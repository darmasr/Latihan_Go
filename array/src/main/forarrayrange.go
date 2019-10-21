package  main 
	
import "fmt"

func main () {

	var buah = [3]string{"apel", "mangga", "anggur"}

	for i, buah := range buah{
		fmt.Printf("elemen %d : %s\n", i, buah)		
	}
}