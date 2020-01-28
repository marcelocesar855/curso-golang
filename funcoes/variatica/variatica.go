package main

import "fmt"

func media(numeros ...float64) float64 { //... representa uma funcao variatica, que recebe n parametros, que é tratado como um array
	total := 0.0
	for _, num := range numeros {
		total += num
	}
	return total / float64(len(numeros))
}

func main() {
	fmt.Printf("Média = %.2f", media(6.6, 5.7, 8.3, 9.8))
}
