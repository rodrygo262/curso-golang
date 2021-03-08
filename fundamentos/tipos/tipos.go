package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	i1 := math.MaxInt64

	fmt.Println("O valor maximo do int é", i1)
	fmt.Println("O tipo de i1 é ", reflect.TypeOf(i1))

	var i2 rune = 'a' // Representa um mapeamento da tabela unicode
	fmt.Println("O rone é", reflect.TypeOf((i2)))
	fmt.Println(i2)

	var x float32 = 49.99
	fmt.Println("O tipo de x é ", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99)) // float64

	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	s1 := "Olá meu nome é Rodrigo"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é ", len(s1))

	s2 := `Olá
	meu
	nome
	é
	Rodrigo`
	fmt.Println(s2)
	fmt.Println("O tamanho de s2 é", len(s2))
}
