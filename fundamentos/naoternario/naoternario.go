package main

import "fmt"

/* em go não existe operador ternario, como
valor := x > y ? "é maior" : "é menor" */

func obterResultado(nota float64) string {
	if nota > 6 {
		return "Aprovado"
	}
	return "Reprovado"
}

func main() {
	fmt.Println(obterResultado(6.5))
}
