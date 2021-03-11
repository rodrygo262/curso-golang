package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	ch <- 1 //Envia os dados para o chanel
	<-ch    //Recebe os dados do chanel

	ch <- 2
	fmt.Println(<-ch)
}
