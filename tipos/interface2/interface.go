package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual int
}

func (f *ferrari) ligarTurbo() { //recebe um ponteiro por alterar valores da variavel
	f.turboLigado = true
}

func main() {
	carro1 := ferrari{"F30", false, 0}
	carro1.ligarTurbo()

	var carro2 esportivo = &ferrari{"F40", false, 0} //quando a veriavel tem uma interface como tipo e um método com ponteiro é necessario fazer a referencia ao endereço de memoria
	carro2.ligarTurbo()

	fmt.Println(carro1, carro2)
}
