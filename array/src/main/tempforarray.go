package  main 
	
import "fmt"

func main () {

	var buah = [3]string{"apel", "mangga", "anggur"}

	for _, buah := range buah {
		fmt.Printf("nama buah : %s\n", buah)
	}
}