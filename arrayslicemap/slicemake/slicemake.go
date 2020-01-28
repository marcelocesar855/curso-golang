package main

import "fmt"

func main() {
	s1 := make([]int, 10) // criando slice do tipo int com tamanho de 10 e array interno
	s1[9] = 12
	fmt.Println(s1)

	s1 = make([]int, 10, 20) // criando slice com array interno de tamanho 20
	fmt.Println(s1, len(s1), cap(s1))

	s1 = append(s1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0) // criando novo slice e atrinuindo valores inclusive ja existentes
	fmt.Println(s1, len(s1), cap(s1))

	s1 = append(s1, 1) // quando o tamanho do array interno é atingido ele tem seu tamanho dobrado
	fmt.Println(s1, len(s1), cap(s1))
}
