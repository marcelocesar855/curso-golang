package main

import "fmt"

func main() {
	letraEnome := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 15000,
			"Gabriel Santos": 6000,
		},
		"H": {
			"Helen Bastos": 1200,
			"Heitor Cesar": 1900,
		},
		"P": {
			"Pedro Jorge": 12000,
		},
	}

	delete(letraEnome, "H")

	for _, nomes := range letraEnome {
		for nome, salario := range nomes {
			fmt.Printf("%s, %.2f\n", nome, salario)
		}
	}
}
