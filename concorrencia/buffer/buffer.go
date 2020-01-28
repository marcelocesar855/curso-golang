package main

import (
	"fmt"
	"time"
)

func rotina(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println("Executou!")
	ch <- 4
	ch <- 5
	ch <- 6
}

func main() {
	ch := make(chan int, 3) //channel com buffer de 3 posicoes
	go rotina(ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch)
}
