package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
}

type bmw7 struct{}

func (b bmw7) ligarTurbo() {
	fmt.Println("Turbo...")
}

func (b bmw7) fazerBaliza() {
	fmt.Println("Uepaaaa baliza...")
}

func main() {
	var esportivoLuxuoso = bmw7{}
	esportivoLuxuoso.fazerBaliza()
	esportivoLuxuoso.ligarTurbo()
}
