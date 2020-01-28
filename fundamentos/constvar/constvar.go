package main

import (
	"fmt"
	m "math"
)

func main() {
	const pi float64 = 3.1415
	var raio = 3.2 //tipo float inferido pelo compilador

	//forma reduzida

	area := pi * m.Pow(raio, 2) // inicializando e atribuindo valor

	fmt.Println("A área da circunferência é", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "opa!"
	fmt.Println(g, h, i)
}
