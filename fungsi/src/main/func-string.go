package main

import "fmt"
import "strings"


func main() {
	var nama = []string{"Darma", "Sukarame"}
	printPesan("halo", nama)
}

func printPesan(pesan string, arr []string) {
	var namaString = strings.Join(arr, " ")
	fmt.Println(pesan, namaString)
}