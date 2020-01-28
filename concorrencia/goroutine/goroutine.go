package main

import (
	"fmt"
	"time"
)

func falar(pessoa, texto string, qntd int) {
	for i := 0; i < qntd; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (interação %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	//falar("Maria", "Porque você não fala comigo?", 3)
	//falar("João", "Só posso falar depois de você", 1)

	//go falar("Maria", "Ei", 10)
	//go falar("João", "Opa", 10)
	//time.Sleep(time.Second * 5)

	go falar("Maria", "Entendi!", 10)
	falar("João", "Parabéns", 5)
}
