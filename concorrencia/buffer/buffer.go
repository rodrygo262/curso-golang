package main

import "fmt"

func rotina(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	fmt.Println("executou")
	ch <- 5
	ch <- 6
}

func main() {
	ch := make(chan int, 3)
	go rotina(ch)

	fmt.Println(<-ch)
}
