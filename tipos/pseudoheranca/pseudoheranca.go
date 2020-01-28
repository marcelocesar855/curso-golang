package main

import "fmt"

type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	carro       // campo anônimo gera composição
	turboLigado bool
}

func (c carro) verificarVelocidade() int {
	return c.velocidadeAtual
}

func main() {
	f := ferrari{
		turboLigado: true,
	}
	f.nome = "F40"
	f.velocidadeAtual = 25

	fmt.Printf("A ferrari %s está com o turbo ligado? %t\n", f.nome, f.turboLigado)

	fmt.Println(f.verificarVelocidade())
}
