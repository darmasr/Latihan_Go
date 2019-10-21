package main

import "fmt"

func main() {
	var ayam = map[string]int{
		"januari": 50,
		"februari": 40,
		"maret": 30,
		"april": 34,
	}

	for key, val := range ayam {
		fmt.Println(key, " \t:", val)
		
	}
}
