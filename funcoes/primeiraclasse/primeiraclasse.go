package main

import "fmt"

var soma = func(a, b int) int { // em go funcoes podem ser armazenadas em variaveis
	return a + b
}

func main() {
	fmt.Println(soma(1, 2))

	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println(sub(2, 1))
}
