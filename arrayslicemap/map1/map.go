package main

import "fmt"

func main() {
	//var aprovados map[int]string
	//maps devem sempre serem inicializados (:=)
	aprovados := make(map[int]string)
	aprovados[12345678900] = "Maria"
	aprovados[12345678901] = "João"
	aprovados[12345678902] = "Ana"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[12345678901]) //apresentando valor de map atraves da chave

	delete(aprovados, 12345678902) //deletando valor através da chave
	fmt.Println(aprovados)
}
