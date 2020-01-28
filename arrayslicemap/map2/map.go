package main

import "fmt"

func main() {
	nomeEsalario := map[string]float64{ //inicializando e atribuindo valores
		"Aroldo": 1200,
		"Mario":  2200,
		"Clara":  3600,
	}

	nomeEsalario["Maria"] = 6000

	delete(nomeEsalario, "Inexistente") //caso não exista a chave não é retornado um erro

	for nome, salario := range nomeEsalario {
		fmt.Printf("%s, %.2f\n", nome, salario)
	}

}
