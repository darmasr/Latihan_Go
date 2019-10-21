package main

import "fmt"

func main() {
	var ayam  map[string]int
	ayam = map[string]int{}

	ayam["januari"] = 50
	ayam["februari"] = 40

	fmt.Println("januari", ayam["januari"])
	fmt.Println("mei", ayam["mei"])
}
