package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	// struct para json
	p1 := produto{1, "Caneta", 1.99, []string{"Papelaria", "Importado"}}
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	// json para struct
	var p2 produto
	jsonString := `{"id":2,"nome":"Notebook","preco":1999.90,"tags":["Informatica", "Promoção"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2.Nome)
}
