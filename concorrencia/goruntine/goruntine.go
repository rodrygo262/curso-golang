package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s ( interação %d )\n", pessoa, texto, i+1)
	}
}

func main() {
	// Roda de forma sequencial
	// fale("Maria", "Pq vc não fala comigo?", 3)
	// fale("João", "Só posso falar depois de você.", 1)

	// Roda de forma de treeds mas precisa que a aplicação principal continue executando para funcionar
	// go fale("Maria", "Eiii", 500)
	// go fale("João", "Oiii", 500)
}
