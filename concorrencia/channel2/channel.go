package main

import (
	"fmt"
	"time"
)

//Channel é uma forma de comunicação entre goroutines
//Channel é um tipo como String e int ('chan int' ou 'chan String' e por ai vai)

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base
}

func main() {
	c := make(chan int)
	go doisTresQuatroVezes(2, c)

	a, b := <-c, <-c //recebendo dados do canal
	fmt.Println(a, b)

	fmt.Println(<-c)
}
