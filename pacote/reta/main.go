package main

import "fmt"

func main() {
	ponto1 := Ponto{2.0, 2.0}
	ponto2 := Ponto{2.0, 4.0}
	fmt.Println(catetos(ponto1, ponto2))
	fmt.Println(Distancia(ponto1, ponto2))
}
