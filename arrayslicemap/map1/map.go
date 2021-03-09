package main

import "fmt"

func main() {
	aprovados := make(map[int]string)
	aprovados[12345678978] = "Maria"
	aprovados[12345678922] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	delete(aprovados, 12345678978)
	fmt.Println(aprovados)
}
