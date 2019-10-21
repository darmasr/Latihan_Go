package  main 
	
import "fmt"

func main () {

	var buah = [3]string{"apel", "mangga", "anggur"}

	for i := 0; i < len(buah); i++ {
		fmt.Printf("elemen %d : %s \n", i, buah[i])
	}
}