package main

import "fmt"

func fatorial(numero uint) uint {
	resultado := numero
	for numero > 1 {
		resultado *= numero - 1
		numero--
	}
	return resultado
}

func main() {
	resultado := fatorial(5)
	fmt.Println(resultado)
}
