package main

import "fmt"

func listaAprovados(alunos ...string) {
	for i, nome := range alunos {
		fmt.Printf("%d) %s\n", i+1, nome)
	}
}

func main() {
	alunos := []string{"José", "Maria", "Carlos", "Ana"} //para passar valores conjuntos para uma funcao variatica necessita ser em um slice
	listaAprovados(alunos...)                            //separa os valores do slice em parametros diferentes para a funcao variatica
}
