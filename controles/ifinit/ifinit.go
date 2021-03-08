package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAreatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	if i := numeroAreatorio(); i > 5 {
		fmt.Println("Ganhou!!")
	} else {
		fmt.Println("Perdeu!")
	}
}
