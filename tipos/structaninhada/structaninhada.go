package main

import "fmt"

type item struct {
	itemID int
	qntd   int
	preco  float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) calcularTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.qntd)
	}
	return total
}

func main() {
	pedido1 := pedido{
		userID: 1,
		itens: []item{
			item{itemID: 1, qntd: 3, preco: 2.99},
			item{2, 5, 7.99},
			item{3, 4, 35.99},
		},
	}
	fmt.Printf("O valor total do pedido é R$ %.2f", pedido1.calcularTotal())
}
