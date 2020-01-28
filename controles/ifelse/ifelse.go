package main

import "fmt"

func obterResultado(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado com nota", nota)
	} else {
		fmt.Println("Reprovado com nota", nota)
	}
}

func main() {
	obterResultado(7.3)
	obterResultado(5.3)
}
