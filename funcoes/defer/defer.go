package main

import "fmt"

func consultarBanco(num int) {
	fmt.Println("Iniciando conexão...")
	defer fmt.Println("Encerrando conexão...")
	if num > 10 {
		fmt.Println("Consultando tabela B...")
		return
	}
	fmt.Println("Consultando tabela A...")
}

func main() {
	consultarBanco(11)
}
