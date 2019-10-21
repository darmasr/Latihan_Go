package main
import "fmt"
 
func main() {
    var baris int
    var n int = 0
    fmt.Print("Masukkan jumlah baris :")
    fmt.Scan(&baris)     
    for i := 1; i <= baris; i++ {     
        n=0
        for space := 1; space <= baris-i; space++ {
            fmt.Print("  ")         
        }
        for {
            fmt.Print("* ")
            n++
            if(n == 2*i-1){             
                break
            }
        }       
        fmt.Println("")
    }
