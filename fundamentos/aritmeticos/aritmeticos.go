package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Soma >", a+b)
	fmt.Println("Subtração >", a-b)
	fmt.Println("Multiplicação >", a*b)
	fmt.Println("Divisão >", a/b)
	fmt.Println("Módulo >", a%b)

	//operações bit a bit
	fmt.Println("E", a&b)
	fmt.Println("Ou", a|b)
	fmt.Println("Xor", a^b)

	//pacote math
	c := 3.0
	d := 2.0
	fmt.Println("Máximo >", math.Max(c, d))
	fmt.Println("Mínimo >", math.Min(c, d))
	fmt.Println("Exponencial", math.Pow(c, d))
}
