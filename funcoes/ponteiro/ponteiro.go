package main

import "fmt"

func int1(n int) { // método que recebe um valor
	n++
}

func int2(n *int) { // método que recebe um ponteiro
	*n++
}

func main() {
	n := 1

	int1(n) // chamada por valor
	fmt.Println(n)

	int2(&n) // chamada por referencia
	fmt.Println(n)
}
