package main

import "fmt"

type item struct {
	produtoID int
	qtde      int
	preco     float64
}

type pedido struct {
	userId int
	itens  []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.qtde)
	}
	return total
}

func main() {
	pedido := pedido{
		userId: 1,
		itens: []item{
			{
				produtoID: 1,
				qtde:      2,
				preco:     12.10,
			},
			{
				produtoID: 2,
				qtde:      1,
				preco:     23.10,
			},
			{
				produtoID: 11,
				qtde:      100,
				preco:     3.10,
			},
		},
	}
	fmt.Printf("O valor total do pedido é %.2f", pedido.valorTotal())
}
