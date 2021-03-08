package main

import "fmt"

func main() {
	i := 1

	var p *int = nil // Pega o endereço da variavel i
	p = &i           // Desreferencia o ponteiro ( pegar o valor )
	*p++
	i++
	fmt.Println(p, *p, i, &i)
}
