package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	Id    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	p1 := produto{
		Id:    1,
		Nome:  "Notebook",
		Preco: 1899.90,
		Tags:  []string{"Promoção", "Eletronico"},
	}

	p1JSon, _ := json.Marshal(p1)

	fmt.Println(string(p1JSon))

	var p2 produto
	jsonString := `{"id": 2, "nome": "Caneta", "preco": 8.90, "tags": ["Papelaria", "Importado"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2.Tags[1])
}
