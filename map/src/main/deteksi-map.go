package main

import "fmt"

func main() {
	var ayam = map[string]int{"januari": 50, "februari": 40}
	var value, isExist = ayam["mei"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item tidak ditemukan")
	}
}
