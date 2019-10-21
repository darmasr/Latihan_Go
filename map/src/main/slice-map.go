package main

import "fmt"

func main() {
	var ayam = []map[string]string{
		map[string]string{"nama": "ayam kampung", "gender": "jantan"},
		map[string]string{"nama": "ayam arab", "gender": "betina"},
		map[string]string{"nama": "ayam cemani", "gender": "betina"},
	}

	for _, ayam1 := range ayam {
		fmt.Println(ayam1["gender"], ayam1["nama"])
	}
}
