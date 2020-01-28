package main

import "fmt"

type nota float64

func (n nota) entre(inicio float64, fim float64) bool {
	return float64(n) >= inicio && float64(n) <= fim
}

func notaParaConceito(n nota) string {
	switch {
	case n.entre(9.0, 10.0):
		return "A"
	case n.entre(7.0, 8.99):
		return "B"
	case n.entre(6.0, 7.99):
		return "C"
	case n.entre(3.0, 5.99):
		return "D"
	default:
		return "E"
	}
}

func main() {
	fmt.Println(notaParaConceito(8.1))
	fmt.Println(notaParaConceito(6.5))
	fmt.Println(notaParaConceito(2.7))
}
