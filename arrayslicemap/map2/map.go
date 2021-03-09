package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"José João":      11325.45,
		"Gabriela Silva": 15455.15,
		"Pedro junior":   1200.00,
	}

	funcsESalarios["Rafael Filho"] = 1350.00
	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
