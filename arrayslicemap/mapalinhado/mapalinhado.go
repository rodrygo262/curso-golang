package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 15482.33,
			"Guga":           1515.33,
		},
		"J": {
			"Jose": 3194.33,
		},
		"Z": {
			"Zumira": 10653.65,
		},
	}

	fmt.Println(funcsPorLetra)

	for letra, funcs := range funcsPorLetra{
		fmt.Println(letra, funcs)
		for nome, salario := range funcs {
			fmt.Println(nome, salario)
		}
	}
}
