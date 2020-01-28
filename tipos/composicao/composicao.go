package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
	pilotoAutomatico()
}

type bmw struct {
	modelo                 string
	turboLigado            bool
	balizaLigada           bool
	pilotoAutomaticoLigado bool
}

func (b *bmw) ligarTurbo() {
	b.turboLigado = true
}

func (b *bmw) fazerBaliza() {
	b.balizaLigada = true
}

func (b *bmw) pilotoAutomatico() {
	b.pilotoAutomaticoLigado = true
}

func main() {
	var carro esportivoLuxuoso = &bmw{"BMW7", false, false, false}
	carro.fazerBaliza()
	carro.ligarTurbo()
	carro.pilotoAutomatico()
	fmt.Println(carro)
}
