package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) {
	tv50 := trab1 && trab2
	tv32 := trab1 != trab2 //operador de Ou excluisivo no go não existe, mas != traz o valor equivalente
	sorvete := trab1 || trab2
	return tv50, tv32, sorvete
}

func main() {
	tv50, tv32, sorvete := compras(false, false)
	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t",
		tv50, tv32, sorvete, !sorvete)
}
