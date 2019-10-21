package  main 
	
import "fmt"

func main () {

	var buah = []string{"apel", "mangga", "anggur"}
	var newBuah = buah[0:2]
	var newBuah1 = buah[1:3]

	fmt.Println(newBuah)
	fmt.Println(newBuah1)

	fmt.Println(len(newBuah))
	fmt.Println(cap(newBuah))

	fmt.Println(len(newBuah1))
	fmt.Println(cap(newBuah1))
}