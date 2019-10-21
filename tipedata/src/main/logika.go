package  main 
	
import "fmt"

func main () {

	var left = false
	var right = true
	
	var leftAndRight = left && right
	fmt.Printf("left && right \t(%t) \n", leftAndRight)

	var lefOrRight = left || right
	fmt.Printf("left || right \t(%t) \n", lefOrRight)

	var leftReverse = !left
	fmt.Printf("!left \t \t(%t) \n", leftReverse)

}