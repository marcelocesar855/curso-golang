package main

import (
	"fmt"

	"github.com/cod3rcursos/html"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

func juntar(entrada1, entrada2 <-chan string) <-chan string { //junta os dados vindos de diferentes canais em um so
	ch := make(chan string)
	go encaminhar(entrada1, ch)
	go encaminhar(entrada2, ch)
	return ch
}

func main() {
	ch := juntar( // consumindo o canal de multiplexao
		html.Titulo("https://www.google.com", "https://www.udemy.com"),
		html.Titulo("https://www.amazon.com", "https://www.youtube.com"),
	)
	fmt.Println(<-ch, " | ", <-ch)
}
