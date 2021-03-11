package main

import (
	"fmt"
	"time"
)

func doisTresQuatroVezes(base int, c chan int) {
	c <- base * 2
	time.Sleep(time.Second)
	c <- base * 3
	time.Sleep(time.Second)
	c <- base * 4
}

func main() {
	c := make(chan int)
	go doisTresQuatroVezes(5, c)

	a, b := <-c, <-c
	fmt.Println(a, b)

	fmt.Println(<-c)
}
