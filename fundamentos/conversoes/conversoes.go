package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// Conversao int to string
	fmt.Println("teste " + strconv.Itoa(1234))

	// Conversao string to int
	num, _ := strconv.Atoi("123")
	fmt.Printf("Numero %d\n", num)

	b, _ := strconv.ParseBool("true")

	if b {
		fmt.Println("Verdadeiro")
	}
}
