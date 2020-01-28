package main

import "fmt"

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

func (p pessoa) toString() string { // método presente na interface
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string { // método presente na interface
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

// em go as interfaces são implementadas implicitamente

func main() {
	var coisa imprimivel = pessoa{"Cleiton", "Ramos"}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	coisa = produto{"Bermuda Jeans", 69.90}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	// uma variavel que tem como tipo uma interface pode receber varios tipos que implementam essa interface (polimorfismo)

	p1 := pessoa{"Kennedy", "Miguel"}
	imprimir(p1)

	// um método que recebe uma interface aceita todos os tipos que implementam a interface
}
