package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", ActionIndex)

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}

func ActionIndex(w http.ResponseWriter, r *http.Request) {
	data := []struct {
		Name string
		Age  int
	}{
		{"Song Khun", 24},
		{"Phum Viphurit", 30},
		{"Nong Lheui", 15},
		{"Cheng Mai", 18},
	}
	jsonInBytes, err := json.Marshal(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(jsonInBytes)
}
