package main

import "fmt"

func main() {
	var notas [3]float64 //homogenea (mesmo tipo) e estática (fixa)
	fmt.Println(notas)

	notas[0], notas[1], notas[2] = 7.1, 5.4, 8.7
	fmt.Println(notas)

	var total float64

	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64(len(notas))
	fmt.Printf("Média %.2f", media)
}
