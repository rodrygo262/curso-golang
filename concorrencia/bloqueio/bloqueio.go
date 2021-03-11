package main

import (
	"fmt"
	"time"
)

func testandoDeadLock(c chan int) {
	time.Sleep(time.Second)
	c <- 1
	fmt.Println("Somente depois de ser lido")
}

func main() {
	c := make(chan int)
	go testandoDeadLock(c)

	fmt.Println(<-c)
	fmt.Println("Depois de ser lido")
	fmt.Println("Vamos gerar o deadlock")
	fmt.Println(<-c)
}
