package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	if i := numeroAleatorio(); i > 5 { //estrutura normalmente suportada apenas no for no go presente tambem no if e no switch
		fmt.Println("Ganhou!")
	} else {
		fmt.Println("Perdeu...")
	}
}
