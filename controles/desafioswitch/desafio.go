package main

import "fmt"

func notaPorConceito(nota float64) string {
	i := int(nota)
	switch {
	case i >= 9:
		return "A"
	case i >= 8:
		return "B"
	case i >= 6:
		return "C"
	case i >= 4:
		return "D"
	case i >= 2:
		return "E"
	case i >= 0:
		return "F"
	default:
		return "Nota Inválida"
	}
}

func main() {
	fmt.Println(notaPorConceito(9.2))
	fmt.Println(notaPorConceito(6.5))
	fmt.Println(notaPorConceito(3.4))
	fmt.Println(notaPorConceito(1.1))
}
