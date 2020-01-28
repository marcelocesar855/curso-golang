package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1 //operacao bloqueante ate o valor ser lido
	fmt.Println("Só depois que foi lido")
}

func main() {
	c := make(chan int) //channel sem buffer

	go rotina(c)

	fmt.Println(<-c) //operação bloqueante ate o valor ser enviado
	fmt.Println("Foi lido")
	fmt.Println(<-c)   //deadlock! pois não há mais valores sendo enviados
	fmt.Println("Fim") //nunca é impresso por causa do deadlock
}
